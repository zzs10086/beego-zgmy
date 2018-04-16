package models

import (
	//"github.com/astaxie/beego"
	//"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

type Article struct {
	Id             int             `orm:"column(id);pk;unique;auto_increment;"`
	Title          string          `orm:"column(title);size(255)"`
	Description    string          `orm:"column(description);size(255)"`
	Source         string          `orm:"column(source);size(30)"`
	Addtime        int             `orm:"column(addtime);size(11)"`
	Catid          int             `orm:"column(catid);size(11)"`
	Thumb          string          `orm:"column(thumb);size(255)"`
	VideoUrl       string          `orm:"column(video_url);size(255)"`
	IsImg          int             `orm:"column(is_img);size(2)"`
	Status         int             `orm:"column(status);size(2)"`
	Click          int             `orm:"column(click);size(2)"`
	Flag           string          `orm:"column(flag);size(30)"`
	Keywords       string          `orm:"column(keywords);size(50)"`
	OriginalUrl    string          `orm:"column(original_url);size(255)"`
	Good           int             `orm:"column(good);size(11)"`
	Bad            int             `orm:"column(bad);size(11)"`
	CreateAt       string          `orm:"column(created_at);"`
	UpdateAt       string          `orm:"column(updated_at);"`
	ArticleContent *ArticleContent `orm:"rel(one);"`
}

type ArticleContent struct {
	Id      int      `orm:"column(id);pk;unique;auto_increment"`
	Aid     int      `orm:"column(aid);size(11);"`
	Content string   `orm:"column(content);"`
	Article *Article `orm:"reverse(one);"` // 设置一对一反向关系(可选)
}

func (this *ArticleContent) TableName() string {
	return "article_content"
}

func init() {

	orm.RegisterDriver("mysql", orm.DRMySQL)
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/mayun?charset=utf8", 30)

	// register model
	//orm.RegisterModel(new(Article))

	orm.RegisterModelWithPrefix("my_", new(Article), new(ArticleContent))

}

/**
*
*获取文章详情
*
 */
func GetArcInfo(aid int) (article Article) {

	article = Article{Id: aid}
	o := orm.NewOrm()
	o.Read(&article, "Id")
	return article
	/*
		err := o.Read(&article)

		if err == orm.ErrNoRows {
			//beego.Info("查询不到")
			return nil
		} else if err == orm.ErrMissPK {
			//beego.Info("找不到主键")
			return nil
		} else {
			//beego.Info("title:", article.Title)
			//return article
		}
	*/
}
