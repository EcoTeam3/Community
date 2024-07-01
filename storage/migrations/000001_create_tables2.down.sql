-- Down migration script

-- Drop Comments table
DROP TABLE IF EXISTS comments;

-- Drop Posts table
DROP TABLE IF EXISTS posts;

-- Drop Group Members table
DROP TABLE IF EXISTS group_members;

-- Drop Groups table
DROP TABLE IF EXISTS groups;

-- Drop Member Role type
DROP TYPE IF EXISTS member_role;
