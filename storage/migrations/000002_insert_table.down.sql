-- Drop all data inserted into comments table in up migration
DELETE FROM comments WHERE comment_id IN (1, 2, 3, 4);

-- Drop all data inserted into posts table in up migration
DELETE FROM posts WHERE post_id IN (1, 2, 3, 4);

-- Drop all data inserted into group_members table in up migration
DELETE FROM group_members WHERE (group_id, user_id) IN (('f7ddcc9f-2924-4b8a-bd0e-7e6d6bb3a0c1', 1), ('f7ddcc9f-2924-4b8a-bd0e-7e6d6bb3a0c1', 2), ('b5a2f1cd-5d4a-4e89-a423-8990b49a4e9f', 3), ('75dc7bca-6e37-4ec8-a518-4b8a0a9b6893', 1), ('75dc7bca-6e37-4ec8-a518-4b8a0a9b6893', 3));

-- Drop all data inserted into groups table in up migration
DELETE FROM groups WHERE group_id IN ('f7ddcc9f-2924-4b8a-bd0e-7e6d6bb3a0c1', 'b5a2f1cd-5d4a-4e89-a423-8990b49a4e9f', '75dc7bca-6e37-4ec8-a518-4b8a0a9b6893');
