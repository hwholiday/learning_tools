package main

import (
	"github.com/urfave/cli/v2"
)

var (
	app             *cli.App
	checkImage      = []string{"JPG", "JPEG", "PNG", "GIF"}
	endpoint        string
	accessKeyID     string
	secretAccessKey string
	useSSL          bool
	maxWidth        uint
	maxHeight       uint
)

func init() {
	app = new(cli.App)
	app.Name = "图片自动生成缩略图服务"
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "endpoint",
			Value: "172.12.15.135:9000",
			Usage: "S3文件服务器地址",
		},
		&cli.StringFlag{
			Name:  "key",
			Value: "minioadmin",
			Usage: "accessKeyID",
		},
		&cli.StringFlag{
			Name:  "access_key",
			Value: "minioadmin",
			Usage: "secretAccessKey",
		},
		&cli.BoolFlag{
			Name:  "ssl",
			Value: false,
			Usage: "是否使用SSl",
		},
		&cli.UintFlag{
			Name:  "width",
			Value: 200,
			Usage: "压缩后图片宽",
		}, &cli.UintFlag{
			Name:  "height",
			Value: 200,
			Usage: "压缩后图片高",
		},
	}
	app.Action = func(c *cli.Context) error {
		endpoint = c.String("endpoint")
		accessKeyID = c.String("key")
		secretAccessKey = c.String("access_key")
		useSSL = c.Bool("ssl")
		maxWidth = c.Uint("width")
		maxHeight = c.Uint("height")
		StartMinio()
		return nil
	}
}
