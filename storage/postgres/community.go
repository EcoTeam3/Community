package postgres

import (
	pb "community/generated"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type NewCommunity struct {
	Db *sql.DB
}

func NewCommunityRepo(db *sql.DB) *NewCommunity {
	return &NewCommunity{Db: db}
}

func (C *NewCommunity) GetGroup(groupId *pb.GroupId) (*pb.Group, error) {
	id, err := uuid.Parse(groupId.GroupId)
	if err != nil {
		return nil, fmt.Errorf("invalid group ID: %w", err)
	}

	group := pb.Group{}
	var userIdBytes []byte
	err = C.Db.QueryRow(`SELECT group_id, name, description, created_by, created_at
       FROM groups
       WHERE group_id = $1 AND deleted_at IS NULL`, id).Scan(&group.GroupId, group.Name, group.Description, group.CreatedBy, group.CreatedAt)
	if err != nil {
		return nil, err
	}

	group.GroupId = &pb.GroupId{GroupId: uuid.UUID(userIdBytes).String()}
	return &group, nil

}


func (C *NewCommunity) DeleteGroup(groupId *pb.GroupId) (*pb.Status, error) {
	id, err := uuid.Parse(groupId.GroupId)
	if err != nil {
		return nil, fmt.Errorf("invalid group ID: %w", err)
	}

	_, err = C.Db.Exec(`UPDATE groups 
		SET deleted_at = $1
		Where group_id = $2`, time.Now(), id)
	if err != nil {
		return &pb.Status{Status: false}, err
	}

	return &pb.Status{Status: true}, nil

}


func (C *NewCommunity) JoinGroupUser(IDs *pb.JoinLeave) (*pb.Status, error) {

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
    `, IDs.GroupId, IDs.UserId, IDs.Role, IDs.JoinedAt)

  if err != nil {
    return &pb.Status{Status: false}, err
  }

  rowsAffected, err := rows.RowsAffected()

  if err != nil || rowsAffected == 0 {
    return &pb.Status{Status: false}, err
  }

  return &pb.Status{Status: true}, err

}