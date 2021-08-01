package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/zhshch2002/goribot"
)

func main() {
	f, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Failed to create file")
	}
	defer f.Close()

	s := goribot.NewSpider(
		goribot.Limiter(true, &goribot.LimitRule{
			Glob: "*.bilibili.com",
			Rate: 2,
		}),
		goribot.RefererFiller(),
		goribot.RandomUserAgent(),
		goribot.SetDepthFirst(true),
	)

	var getVideoInfo = func(ctx *goribot.Context) {
		res := map[string]interface{}{
			"bvid":  ctx.Resp.Json("data.bvid").String(),
			"title": ctx.Resp.Json("data.title").String(),
			"des":   ctx.Resp.Json("data.des").String(),
			"pic":   ctx.Resp.Json("data.pic").String(),
			"tname": ctx.Resp.Json("data.tname").String(),
			"owner": map[string]interface{}{
				"name": ctx.Resp.Json("data.owner.name").String(),
				"mid":  ctx.Resp.Json("data.owner.mid").String(),
				"face": ctx.Resp.Json("data.owner.face").String(),
			},
			"ctime":   ctx.Resp.Json("data.ctime").String(),
			"pubdate": ctx.Resp.Json("data.pubdate").String(),
			"stat": map[string]interface{}{
				"view":     ctx.Resp.Json("data.stat.view").Int(),
				"danmaku":  ctx.Resp.Json("data.stat.danmaku").Int(),
				"reply":    ctx.Resp.Json("data.stat.reply").Int(),
				"favorite": ctx.Resp.Json("data.stat.favorite").Int(),
				"coin":     ctx.Resp.Json("data.stat.coin").Int(),
				"share":    ctx.Resp.Json("data.stat.share").Int(),
				"like":     ctx.Resp.Json("data.stat.like").Int(),
				"dislike":  ctx.Resp.Json("data.stat.dislike").Int(),
			},
		}
		ctx.AddItem(res)
	}

	var findVideo goribot.CtxHandlerFun
	findVideo = func(ctx *goribot.Context) {
		u := ctx.Req.URL.String()
		fmt.Println(u)
		if strings.HasPrefix(u, "https://www.bilibili.com/video/") {
			if strings.Contains(u, "?") {
				u = u[:strings.Index(u, "?")]
			}
			u = u[31:]
			fmt.Println(u)

			ctx.AddTask(goribot.GetReq("https://api.bilibili.com/x/web-interface/view?bvid="+u), getVideoInfo)
		}
		ctx.Resp.Dom.Find("a[href]").Each(func(i int, sel *goquery.Selection) {
			if h, ok := sel.Attr("href"); ok {
				ctx.AddTask(goribot.GetReq(h), findVideo)
			}
		})
	}

	s.OnItem(func(i interface{}) interface{} {
		data, ok := i.(map[string]interface{})
		if !ok {
			fmt.Println("Cannot cast to map")
		}
		output := fmt.Sprintf("Title: %s Views: %d\n", data["title"], data["stat"].(map[string]interface{})["view"])
		fmt.Println(output)
		f.WriteString(output)
		f.Sync()
		return i
	})

	s.AddTask(goribot.GetReq("https://www.bilibili.com/video/BV1j44y1m76B"), findVideo)
	s.Run()
}
