package postgres

import (
	pb "community/generated/community"
	"community/generated/user"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type NewCommunity struct {
	Db          *sql.DB
	UserService user.UserServiceClient
}

func NewCommunityRepo(db *sql.DB) *NewCommunity {
	return &NewCommunity{Db: db}
}

func (C *NewCommunity) CreateGroup(group *pb.Group) (*pb.Status, error) {
	_, err := C.Db.Exec(`INSERT INTO
							Groups(name, description, created_by)
						VALUES($1, $2, $3)`,
		group.Name, group.Description, group.CreatedBy)
	if err != nil {
		return &pb.Status{Status: false}, err
	}
	return &pb.Status{Status: true}, nil
}

func (C *NewCommunity) GetGroup(groupId *pb.GroupId) (*pb.Group, error) {

	group := pb.Group{}
	err := C.Db.QueryRow(`SELECT group_id, name, description, created_by, created_at
       FROM groups
       WHERE group_id = $1 `, groupId.GroupId).Scan(&group.GroupId, &group.Name, &group.Description, &group.CreatedBy, &group.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &group, nil
}

func (C *NewCommunity) UpdateGroup(group *pb.Group) (*pb.Status, error) { // TODO

	arr := []interface{}{group.GroupId}
	var param []string
	query := "UPDATE Groups SET group_id = $1"
	if len(group.Name) > 0 {
		arr = append(arr, group.Name)
		query += ", name = :name"
		param = append(param, ":name")
	}
	if len(group.Description) > 0 {
		arr = append(arr, group.Description)
		query += ", description = :description"
		param = append(param, ":description")
	}
	if len(group.CreatedBy) > 0 {
		arr = append(arr, group.CreatedBy)
		query += ", created_by = :created_by"
		param = append(param, ":created_by")
	}
	n := 1
	for _, j := range param {
		query = strings.Replace(query, j, "$"+strconv.Itoa(n), 1)
		n++
	}
	query += fmt.Sprintf(" WHERE group_id = $%d", n)
	arr = append(arr, group.GroupId)

	_, err := C.Db.Exec(query, arr...)
	if err != nil {
		return &pb.Status{Status: false}, err
	}
	return &pb.Status{Status: true}, nil
}

func (C *NewCommunity) DeleteGroup(groupId *pb.GroupId) (*pb.Status, error) {

	_, err := C.Db.Exec(`delete from groups where group_id = $1`, groupId.GroupId)
	if err != nil {
		return &pb.Status{Status: false}, err
	}

	return &pb.Status{Status: true}, nil

}

func (C *NewCommunity) GetAllGroups(req *pb.Req) (*pb.Groups, error) {
	rows, err := C.Db.Query(`SELECT * FROM Groups`)
	if err != nil {
		return nil, err
	}
	groups := pb.Groups{}
	for rows.Next() {
		var group = pb.Group{}
		err = rows.Scan(&group.GroupId, &group.Name, &group.Description, &group.CreatedBy, &group.CreatedAt)
		if err != nil {
			return nil, err

		}
		groups.Groups = append(groups.Groups, &group)
	}
	return &groups, nil
}

func (C *NewCommunity) JoinGroupUser(IDs *pb.JoinLeave) (*pb.Status, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	user, err := C.UserService.GetUser(ctx, &user.UserId{UserId: IDs.UserId})
	if user == nil || err != nil {
		return &pb.Status{Status: false}, errors.New("user not found")
	}
	group, err := C.GetGroup(&pb.GroupId{GroupId: IDs.GroupId})
	if group == nil || err != nil {
		return &pb.Status{Status: false}, errors.New("group not found")
	}
	rows, err := C.Db.Exec(`
    INSERT 
    INTO 
    group_members(
      group_id,
      user_id,
      role,
      joined_at
    )
    VALUES(
      $1,
      $2,
      $3,
      $4)
    `, IDs.GroupId, IDs.UserId, IDs.Role, time.Now())

	if err != nil {
		return &pb.Status{Status: false}, err
	}

	rowsAffected, err := rows.RowsAffected()

	if err != nil || rowsAffected == 0 {
		return &pb.Status{Status: false}, err
	}
	return &pb.Status{Status: true}, err
}

func (C *NewCommunity) LeaveGroupUser(joinLeave *pb.JoinLeave) (*pb.Status, error) {
	_, err := C.Db.Exec("Update Group_Members SET deleted_at = $1 WHERE group_id = $2 and user_id = $3", time.Now(), joinLeave.GroupId, joinLeave.UserId)
	if err != nil {
		return &pb.Status{
			Status: false,
		}, err
	}
	return &pb.Status{
		Status: true,
	}, err
}

func (C *NewCommunity) UpdateGroupMeber(userRole *pb.UserRole) (*pb.Status, error) {
	_, err := C.Db.Exec("UPDATE Group_Members SET Role = $1 WHERE group_id = $2 AND user_id = $3", userRole.Role, userRole.GroupId, userRole.UserId)
	if err != nil {
		return &pb.Status{
			Status: false,
		}, err
	}
	return &pb.Status{
		Status: true,
	}, nil
}


func (C *NewCommunity) CreatePost(post *pb.Post) (*pb.Status, error) {
	_, err := C.Db.Exec("INSERT INTO Posts(group_id,user_id,content,created_at) VALUES($1,$2,$3,$4)",
		post.GroupId, post.UserId, post.Content, time.Now())
	if err != nil {
		return &pb.Status{
			Status: false,
		}, err
	}
	return &pb.Status{
		Status: true,
	}, nil
}

func (C *NewCommunity) UpdatePost(post *pb.Post) (*pb.Status, error) {
	arr := []interface{}{post.PostId}
	var param []string
	query := "UPDATE Posts SET post_id = $1 "

	if len(post.GroupId) > 0 {
		arr = append(arr, post.GroupId)
		query += ", group_id = :group_id "
		param = append(param, ":group_id")
	}
	if len(post.UserId) > 0 {
		arr = append(arr, post.UserId)
		query += ", user_id = :user_id"
		param = append(param, ":user_id")
	}
	if len(post.Content) > 0 {
		arr = append(arr, post.Content)
		query += ", content = :content "
		param = append(param, ":content")
	}
	if len(post.CreatedAt) > 0 {
		arr = append(arr, post.CreatedAt)
		query += ", created_at = :created_at "
		param = append(param, ":created_at")
	}
	n := 1
	for _, j := range param {
		query = strings.Replace(query, j, "$"+strconv.Itoa(n), 1)
		n++
	}
	query = fmt.Sprintf(" WHERE post_id = $%d", n)
	arr = append(arr, post.PostId)

	_, err := C.Db.Exec(query, arr...)

	if err != nil {
		return &pb.Status{
			Status: false,
		}, err
	}
	return &pb.Status{
		Status: true,
	}, nil
}

func (C *NewCommunity) DeletePost(postId *pb.PostId) (*pb.Status, error) {
	_, err := C.Db.Exec("DELETE FROM Posts WHERE post_id = $1", postId)
	if err != nil {
		return &pb.Status{
			Status: false,
		}, err
	}
	return &pb.Status{
		Status: true,
	}, nil
}

func (C *NewCommunity) GetPost(postId *pb.PostId) (*pb.Post, error) {
	post := pb.Post{}
	err := C.Db.QueryRow("SELECT * FROM Posts WHERE post_id = $1", postId).Scan(
		&post.PostId, &post.GroupId, &post.UserId, &post.Content, &post.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (C *NewCommunity) GetGroupPost(groupPost *pb.GroupPost) (*pb.Post, error) {
	post := pb.Post{}
	err := C.Db.QueryRow("SELECT * FROM Posts WHERE post_id = $1 AND group_id = $2", groupPost.PostId, groupPost.GroupId).Scan(
		post.PostId, post.GroupId, post.UserId, post.Content, post.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (C *NewCommunity) CreatePostComments(comment *pb.Comment) (*pb.Status, error) {
	_, err := C.Db.Exec(`INSERT INTO Comments(comment_id, post_id, user_id, content) VALUES($1, $2, $3, $4)`,
		comment.CommentId, comment.PostId, comment.UserId, comment.Content)
	if err != nil {
		return &pb.Status{Status: false}, err
	}
	return &pb.Status{Status: true}, nil
}

func (C *NewCommunity) GetPostComments(postComment *pb.PostComment) (*pb.Comment, error) {
	comment := &pb.Comment{}
	err := C.Db.QueryRow(`SELECT comment_id, post_id, user_id, content
							FROM
						Comments
							WHERE comment_id = $1, post_id = $2`,
		postComment.CommentId, postComment.PostId).Scan(
		&comment.CommentId, &comment.PostId, &comment.PostId, &comment.Content)
	return comment, err
}
