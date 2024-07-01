-- Delete sample data from Comments table
DELETE FROM comments
WHERE content IN ('Comment content 1', 'Comment content 2', 'Comment content 3');

-- Delete sample data from Posts table
DELETE FROM posts
WHERE content IN ('Post content 1', 'Post content 2', 'Post content 3');

-- Delete sample data from Group Members table
DELETE FROM group_members
WHERE user_id IN ((SELECT user_id FROM users WHERE username='user1'), 
                  (SELECT user_id FROM users WHERE username='user2'), 
                  (SELECT user_id FROM users WHERE username='user3'));

-- Delete sample data from Groups table
DELETE FROM groups
WHERE name IN ('Group 1', 'Group 2');

-- Delete sample data from Users table
DELETE FROM users
WHERE username IN ('user1', 'user2', 'user3');
