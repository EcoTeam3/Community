-- insert_table_up.sql

-- Insert initial data into groups table
INSERT INTO groups (name, description, created_by)
VALUES
('Tech Enthusiasts', 'A group for people who love technology.', gen_random_uuid()),
('Book Club', 'A group for book lovers.', gen_random_uuid()),
('Fitness Freaks', 'A group for fitness enthusiasts.', gen_random_uuid());

-- Insert initial data into group_members table
INSERT INTO group_members (group_id, user_id, role)
VALUES
((SELECT group_id FROM groups WHERE name='Tech Enthusiasts'), gen_random_uuid(), 'admin'),
((SELECT group_id FROM groups WHERE name='Tech Enthusiasts'), gen_random_uuid(), 'member'),
((SELECT group_id FROM groups WHERE name='Book Club'), gen_random_uuid(), 'admin'),
((SELECT group_id FROM groups WHERE name='Book Club'), gen_random_uuid(), 'member'),
((SELECT group_id FROM groups WHERE name='Fitness Freaks'), gen_random_uuid(), 'admin'),
((SELECT group_id FROM groups WHERE name='Fitness Freaks'), gen_random_uuid(), 'member');

-- Insert initial data into posts table
INSERT INTO posts (group_id, user_id, content)
VALUES
((SELECT group_id FROM groups WHERE name='Tech Enthusiasts'), gen_random_uuid(), 'What do you think about the new AI technology?'),
((SELECT group_id FROM groups WHERE name='Book Club'), gen_random_uuid(), 'Has anyone read the latest thriller by John Grisham?'),
((SELECT group_id FROM groups WHERE name='Fitness Freaks'), gen_random_uuid(), 'Share your favorite workout routines.');

-- Insert initial data into comments table
INSERT INTO comments (post_id, user_id, content)
VALUES
((SELECT post_id FROM posts WHERE content='What do you think about the new AI technology?'), gen_random_uuid(), 'I think it''s revolutionary!'),
((SELECT post_id FROM posts WHERE content='What do you think about the new AI technology?'), gen_random_uuid(), 'AI will change the world as we know it.'),
((SELECT post_id FROM posts WHERE content='Has anyone read the latest thriller by John Grisham?'), gen_random_uuid(), 'Yes, it was a real page-turner!'),
((SELECT post_id FROM posts WHERE content='Share your favorite workout routines.'), gen_random_uuid(), 'I love doing HIIT workouts.');
