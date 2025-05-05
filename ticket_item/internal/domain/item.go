package domain

type Item struct {
	Id              int64  `json:"id,omitempty"`
	ItemName        string `json:"itemName,omitempty"`        //活动名称
	AbstractMessage string `json:"abstractMessage,omitempty"` //活动摘要
	ItemTypeId      int64  `json:"itemTypeId,omitempty"`      //活动类型id，如论坛啊，通告啊什么的
	//ItemTypeName        int64  `json:"itemTypName,omitempty"`         //活动类型名称，如论坛啊，通告啊什么的
	ScopeId             int64  `json:"scopeId,omitempty"`             //公共活动、校内活动、校友活动等，以后可 能还有别的范围
	BasicDescription    string `json:"basicDescription,omitempty"`    //基本描述
	ProjectDescription  string `json:"projectDescription,omitempty"`  //详细描述
	ReminderDescription string `json:"reminderDescription,omitempty"` //温馨提示
	State               int8   `json:"state,omitempty"`               //活动状态
	PlaceId             int64  `json:"placeId,omitempty"`             //场馆Id
	CommentCount        int64  `json:"commentCount,omitempty"`        //总评论数
	StartTime           int64  `json:"startTime,omitempty"`           //活动开始日期
	EndTime             int64  `json:"endTime,omitempty"`             //活动结束日期
}
