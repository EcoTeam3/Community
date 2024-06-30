-- Insert initial data into groups table
INSERT INTO groups (group_id, name, description, created_by, created_at)
VALUES 
    ('f7ddcc9f-2924-4b8a-bd0e-7e6d6bb3a0c1', 'Group 1', 'Description of Group 1', 1, CURRENT_TIMESTAMP),
    ('b5a2f1cd-5d4a-4e89-a423-8990b49a4e9f', 'Group 2', 'Description of Group 2', 2, CURRENT_TIMESTAMP),
    ('75dc7bca-6e37-4ec8-a518-4b8a0a9b6893', 'Group 3', 'Description of Group 3', 1, CURRENT_TIMESTAMP);

-- Insert initial data into group_members table
INSERT INTO group_members (group_id, user_id, role, joined_at)
VALUES
    ('f7ddcc9f-2924-4b8a-bd0e-7e6d6bb3a0c1', 1, 'admin', CURRENT_TIMESTAMP),
    ('f7ddcc9f-2924-4b8a-bd0e-7e6d6bb3a0c1', 2, 'member', CURRENT_TIMESTAMP),
    ('b5a2f1cd-5d4a-4e89-a423-8990b49a4e9f', 3, 'moderator', CURRENT_TIMESTAMP),
    ('75dc7bca-6e37-4ec8-a518-4b8a0a9b6893', 1, 'member', CURRENT_TIMESTAMP),
    ('75dc7bca-6e37-4ec8-a518-4b8a0a9b6893', 3, 'admin', CURRENT_TIMESTAMP);

-- Insert initial data into posts table
INSERT INTO posts (post_id, group_id, user_id, content, created_at)
VALUES
    (1, 'f7ddcc9f-2924-4b8a-bd0e-7e6d6bb3a0c1', 1, 'Post 1 in Group 1', CURRENT_TIMESTAMP),
    (2, 'f7ddcc9f-2924-4b8a-bd0e-7e6d6bb3a0c1', 2, 'Post 2 in Group 1', CURRENT_TIMESTAMP),
    (3, 'b5a2f1cd-5d4a-4e89-a423-8990b49a4e9f', 3, 'Post 1 in Group 2', CURRENT_TIMESTAMP),
    (4, '75dc7bca-6e37-4ec8-a518-4b8a0a9b6893', 1, 'Post 1 in Group 3', CURRENT_TIMESTAMP);

-- Insert initial data into comments table
INSERT INTO comments (comment_id, post_id, user_id, content, created_at)
VALUES
    (1, 1, 1, 'Comment 1 on Post 1', CURRENT_TIMESTAMP),
    (2, 1, 2, 'Comment 2 on Post 1', CURRENT_TIMESTAMP),
    (3, 2, 3, 'Comment 1 on Post 2', CURRENT_TIMESTAMP),
    (4, 3, 1, 'Comment 1 on Post 1 in Group 2', CURRENT_TIMESTAMP);
