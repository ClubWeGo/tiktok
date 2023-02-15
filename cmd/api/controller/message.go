package controller

import (
	"context"
	"fmt"
	"github.com/a76yyyy/tiktok/cmd/api2/rpc"
	"github.com/a76yyyy/tiktok/dal/pack"
	"github.com/a76yyyy/tiktok/kitex_gen/message"
	"github.com/a76yyyy/tiktok/pkg/errno"
	"github.com/gin-gonic/gin"
	"strconv"
)

var tempChat = map[string][]Message{}

type ChatResponse struct {
	Response
	MessageList []Message `json:"message_list"`
}

// MessageAction no practical effect, just check if token is valid
func MessageAction(c *gin.Context) {
	fmt.Printf("发送消息")
	var paramVar MessageActionParam
	token := c.Query("token")
	toUserId := c.Query("to_user_id")
	content := c.Query("content")
	action_type := c.Query("action_type")

	vid, err := strconv.Atoi(toUserId)
	if err != nil {
		SendResponse(c, pack.BuildMessageActionResp(errno.ErrBind))
		return
	}
	act, err := strconv.Atoi(action_type)
	if err != nil {
		SendResponse(c, pack.BuildMessageActionResp(errno.ErrBind))
		return
	}

	paramVar.Token = token
	paramVar.ToUserId = int64(vid)
	paramVar.ActionType = int32(act)
	paramVar.Content = &content

	rpcReq := message.DouyinMessageActionRequest{
		Token:      paramVar.Token,
		ToUserId:   paramVar.ToUserId,
		ActionType: paramVar.ActionType,
		Content:    *paramVar.Content,
	}

	ctx := context.Background()
	resp, err := rpc.MessageAction(ctx, &rpcReq)
	if err != nil {
		SendResponse(c, pack.BuildMessageActionResp(errno.ConvertErr(err)))
		return
	}
	SendResponse(c, resp)
}

// MessageChat all users have same follow list
func MessageChat(c *gin.Context) {
	fmt.Printf("获取消息记录\n")
	var paramVar MessageChatParam
	token := c.Query("token")
	toUserId := c.Query("to_user_id")

	tuid, err := strconv.Atoi(toUserId)
	if err != nil {
		SendResponse(c, pack.BuildMessageChatResp(errno.ErrBind))
		return
	}
	paramVar.ToUserId = int64(tuid)
	paramVar.Token = token

	if len(paramVar.Token) == 0 || paramVar.ToUserId < 0 {
		SendResponse(c, pack.BuildMessageChatResp(errno.ErrBind))
		return
	}
	ctx := context.Background()
	resp, err := rpc.MessageChat(ctx, &message.DouyinMessageChatRequest{
		Token:    paramVar.Token,
		ToUserId: paramVar.ToUserId,
	})
	if err != nil {
		SendResponse(c, pack.BuildMessageChatResp(errno.ConvertErr(err)))
		return
	}
	fmt.Printf("成功获取消息记录,%v\n", resp)

	//messages := make([]Message, 0)
	//messages = append(messages, Message{
	//	Id:         1,
	//	Content:    "ddd",
	//	CreateTime: "2-7",
	//})
	//for _, v := range resp.MessageList {
	//	messages = append(messages, Message{
	//		Id:         v.Id,
	//		Content:    v.Content,
	//		CreateTime: v.CreateTime,
	//	})
	//}

	//curMessage := Message{
	//	Id:         1,
	//	FromUserId: 1,
	//	ToUserId:   2,
	//	Content:    "aaa",
	//	CreateTime: "2-7",
	//}
	//chatKey := "1"
	//if messages, exist := tempChat[chatKey]; exist {
	//	tempChat[chatKey] = append(messages, curMessage)
	//} else {
	//	tempChat[chatKey] = []Message{curMessage}
	//}
	//c.JSON(http.StatusOK, ChatResponse{Response: Response{StatusCode: 0}, MessageList: tempChat[chatKey]})
	SendResponse(c, resp)
}
