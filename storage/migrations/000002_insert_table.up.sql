-- Insert sample data into Users table
INSERT INTO users (user_id, username, email, password_hash, created_at) VALUES 
(uuid_generate_v4(), 'user1', 'user1@example.com', 'password1', CURRENT_TIMESTAMP),
(uuid_generate_v4(), 'user2', 'user2@example.com', 'password2', CURRENT_TIMESTAMP),
(uuid_generate_v4(), 'user3', 'user3@example.com', 'password3', CURRENT_TIMESTAMP);

-- Insert sample data into Groups table
INSERT INTO groups (group_id, name, description, created_by, created_at) VALUES
(uuid_generate_v4(), 'Group 1', 'Description for Group 1', (SELECT user_id FROM users WHERE username='user1'), CURRENT_TIMESTAMP),
(uuid_generate_v4(), 'Group 2', 'Description for Group 2', (SELECT user_id FROM users WHERE username='user2'), CURRENT_TIMESTAMP);

-- Insert sample data into Group Members table
INSERT INTO group_members (group_id, user_id, role, joined_at) VALUES
((SELECT group_id FROM groups WHERE name='Group 1'), (SELECT user_id FROM users WHERE username='user1'), 'admin', CURRENT_TIMESTAMP),
((SELECT group_id FROM groups WHERE name='Group 1'), (SELECT user_id FROM users WHERE username='user2'), 'moderator', CURRENT_TIMESTAMP),
((SELECT group_id FROM groups WHERE name='Group 2'), (SELECT user_id FROM users WHERE username='user3'), 'member', CURRENT_TIMESTAMP);

-- Insert sample data into Posts table
INSERT INTO posts (post_id, group_id, user_id, content, created_at) VALUES
(uuid_generate_v4(), (SELECT group_id FROM groups WHERE name='Group 1'), (SELECT user_id FROM users WHERE username='user1'), 'Post content 1', CURRENT_TIMESTAMP),
(uuid_generate_v4(), (SELECT group_id FROM groups WHERE name='Group 1'), (SELECT user_id FROM users WHERE username='user2'), 'Post content 2', CURRENT_TIMESTAMP),
(uuid_generate_v4(), (SELECT group_id FROM groups WHERE name='Group 2'), (SELECT user_id FROM users WHERE username='user3'), 'Post content 3', CURRENT_TIMESTAMP);

-- Insert sample data into Comments table
INSERT INTO comments (comment_id, post_id, user_id, content, created_at) VALUES
(uuid_generate_v4(), (SELECT post_id FROM posts WHERE content='Post content 1'), (SELECT user_id FROM users WHERE username='user2'), 'Comment content 1', CURRENT_TIMESTAMP),
(uuid_generate_v4(), (SELECT post_id FROM posts WHERE content='Post content 2'), (SELECT user_id FROM users WHERE username='user1'), 'Comment content 2', CURRENT_TIMESTAMP),
(uuid_generate_v4(), (SELECT post_id FROM posts WHERE content='Post content 3'), (SELECT user_id FROM users WHERE username='user1'), 'Comment content 3', CURRENT_TIMESTAMP);
