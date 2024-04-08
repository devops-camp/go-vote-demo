package logic

import (
	"github.com/devops-camp/go-vote-demo/app/model"
	"github.com/devops-camp/go-vote-demo/app/tools"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetVoteHandler(c *gin.Context) {
	idStr := c.Query("id")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, tools.EcodeBadRequest("id is required"))

		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, tools.EcodeBadRequest("id is invalid"))
		return
	}

	vote, err := model.GetVote(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, tools.EcodeBadRequest(err.Error()))
		return
	}

	opts, err := model.GetVoteOptsByVoteId(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, tools.EcodeBadRequest(err.Error()))
		return
	}

	data := map[string]any{
		"Vote": vote,
		"Opts": opts,
	}
	//c.JSON(http.StatusOK, vote)
	c.HTML(http.StatusOK, "vote.html", data)
}

type PostVoteParams struct {
	VoteId int64   `form:"vote_id" json:"vote_id" binding:"required"`
	Opts   []int64 `form:"opts" json:"opts" binding:"required"`
}

func PostVoteHandler(c *gin.Context) {
	p := &PostVoteParams{}
	if err := c.ShouldBind(p); err != nil {
		c.JSON(http.StatusBadRequest, tools.EcodeBadRequest(err.Error()))
		return
	}

	for _, id := range p.Opts {
		err := model.UpdateVoteCount(id, p.VoteId)
		if err != nil {
			panic(err)
		}
	}
	//c.JSON(http.StatusOK, p)
	//c.JSON(http.StatusOK, tools.EcodeSuccess("success"))
	c.Redirect(http.StatusSeeOther, "/vote?id="+strconv.FormatInt(p.VoteId, 10))
}
