-- Groups table
CREATE TABLE groups (
    group_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    description TEXT,
    created_by INTEGER REFERENCES users(user_id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TYPE member_role AS ENUM ('member', 'moderator', 'admin');

CREATE TABLE group_members (
    group_id INTEGER REFERENCES groups(group_id),
    user_id INTEGER REFERENCES users(user_id),
    role member_role DEFAULT 'member',
    joined_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (group_id, user_id)
);

-- Posts table
CREATE TABLE posts (
    post_id SERIAL PRIMARY KEY,
    group_id INTEGER REFERENCES groups(group_id),
    user_id INTEGER REFERENCES users(user_id),
    content TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Comments table
CREATE TABLE comments (
    comment_id SERIAL PRIMARY KEY,
    post_id INTEGER REFERENCES posts(post_id),
    user_id INTEGER REFERENCES users(user_id),
    content TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);