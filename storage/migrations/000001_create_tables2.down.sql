-- Drop foreign key constraints from comments table
ALTER TABLE comments DROP CONSTRAINT comments_post_id_fkey;
ALTER TABLE comments DROP CONSTRAINT comments_user_id_fkey;

-- Drop foreign key constraints from posts table
ALTER TABLE posts DROP CONSTRAINT posts_group_id_fkey;
ALTER TABLE posts DROP CONSTRAINT posts_user_id_fkey;

-- Drop foreign key constraints from group_members table
ALTER TABLE group_members DROP CONSTRAINT group_members_group_id_fkey;
ALTER TABLE group_members DROP CONSTRAINT group_members_user_id_fkey;

-- Drop foreign key constraints from groups table
ALTER TABLE groups DROP CONSTRAINT groups_created_by_fkey;

-- Drop tables in reverse order of creation

DROP TABLE IF EXISTS comments;
DROP TABLE IF EXISTS posts;
DROP TABLE IF EXISTS group_members;
DROP TABLE IF EXISTS groups;

-- Drop enum type
DROP TYPE IF EXISTS member_role;
