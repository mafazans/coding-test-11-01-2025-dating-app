CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE profiles (
    id SERIAL PRIMARY KEY,
    user_id INTEGER UNIQUE REFERENCES users(id),
    name VARCHAR(255) NOT NULL,
    bio TEXT,
    age INTEGER,
    gender VARCHAR(50),
    photo_url VARCHAR(255),
    is_verified BOOLEAN DEFAULT false,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE swipes (
    id SERIAL PRIMARY KEY,
    swiper_id INTEGER REFERENCES users(id),
    swiped_id INTEGER REFERENCES users(id),
    is_like BOOLEAN NOT NULL, -- true for right swipe (like), false for left swipe (dislike)
    UNIQUE(swiper_id, swiped_id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX idx_swipes_swiper_date ON swipes(swiper_id, created_at);

CREATE TYPE subscription_status AS ENUM ('active', 'expired', 'cancelled');

CREATE TABLE subscriptions (
    id SERIAL PRIMARY KEY,
    user_id INTEGER UNIQUE REFERENCES users(id),
    status subscription_status NOT NULL DEFAULT 'active',
    start_date TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    end_date TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);
CREATE INDEX idx_subscriptions_user ON subscriptions(user_id);
CREATE INDEX idx_subscriptions_status ON subscriptions(status);

INSERT INTO users (username, password) VALUES
('alice.123@gmail.com', '$2a$10$dummyhashpassword123456'), -- Using a dummy bcrypt hash
('bob.456@yahoo.com', '$2a$10$dummyhashpassword123456'),
('charlie.789@hotmail.com', '$2a$10$dummyhashpassword123456'),
('diana.234@outlook.com', '$2a$10$dummyhashpassword123456'),
('edward.567@gmail.com', '$2a$10$dummyhashpassword123456'),
('fiona.890@yahoo.com', '$2a$10$dummyhashpassword123456'),
('george.345@hotmail.com', '$2a$10$dummyhashpassword123456'),
('hannah.678@outlook.com', '$2a$10$dummyhashpassword123456'),
('ian.901@gmail.com', '$2a$10$dummyhashpassword123456'),
('julia.234@yahoo.com', '$2a$10$dummyhashpassword123456'),
('alice.567@hotmail.com', '$2a$10$dummyhashpassword123456'),
('bob.890@outlook.com', '$2a$10$dummyhashpassword123456'),
('charlie.123@gmail.com', '$2a$10$dummyhashpassword123456'),
('diana.456@yahoo.com', '$2a$10$dummyhashpassword123456'),
('edward.789@hotmail.com', '$2a$10$dummyhashpassword123456'),
('fiona.234@outlook.com', '$2a$10$dummyhashpassword123456'),
('george.567@gmail.com', '$2a$10$dummyhashpassword123456'),
('hannah.890@yahoo.com', '$2a$10$dummyhashpassword123456'),
('ian.123@hotmail.com', '$2a$10$dummyhashpassword123456'),
('julia.456@outlook.com', '$2a$10$dummyhashpassword123456'),
('test1@gmail.com', '$2a$10$IBwaJDJBRZopqemFrCyL7uPe8qqSNg9RzdFsintvXkZz5ohDJH/QK');

INSERT INTO profiles (user_id, name, bio, age, gender, photo_url) VALUES
(1, 'Alice-xyz12', 'Hi! I''m Alice-xyz12 and I love hiking and reading', 25, 'female', 'https://placeholder.com/user1.jpg'),
(2, 'Bob-abc34', 'Hi! I''m Bob-abc34 and I love photography and cooking', 30, 'male', 'https://placeholder.com/user2.jpg'),
(3, 'Charlie-def56', 'Hi! I''m Charlie-def56 and I love traveling and music', 28, 'male', 'https://placeholder.com/user3.jpg'),
(4, 'Diana-ghi78', 'Hi! I''m Diana-ghi78 and I love art and sports', 35, 'female', 'https://placeholder.com/user4.jpg'),
(5, 'Edward-jkl90', 'Hi! I''m Edward-jkl90 and I love gaming and dancing', 27, 'male', 'https://placeholder.com/user5.jpg'),
(6, 'Fiona-mno12', 'Hi! I''m Fiona-mno12 and I love reading and hiking', 32, 'female', 'https://placeholder.com/user6.jpg'),
(7, 'George-pqr34', 'Hi! I''m George-pqr34 and I love cooking and photography', 29, 'male', 'https://placeholder.com/user7.jpg'),
(8, 'Hannah-stu56', 'Hi! I''m Hannah-stu56 and I love music and traveling', 31, 'female', 'https://placeholder.com/user8.jpg'),
(9, 'Ian-vwx78', 'Hi! I''m Ian-vwx78 and I love sports and art', 33, 'male', 'https://placeholder.com/user9.jpg'),
(10, 'Julia-yza90', 'Hi! I''m Julia-yza90 and I love dancing and gaming', 26, 'female', 'https://placeholder.com/user10.jpg'),
(11, 'Alice-bcd12', 'Hi! I''m Alice-bcd12 and I love hiking and music', 28, 'female', 'https://placeholder.com/user11.jpg'),
(12, 'Bob-efg34', 'Hi! I''m Bob-efg34 and I love photography and sports', 34, 'male', 'https://placeholder.com/user12.jpg'),
(13, 'Charlie-hij56', 'Hi! I''m Charlie-hij56 and I love cooking and art', 29, 'male', 'https://placeholder.com/user13.jpg'),
(14, 'Diana-klm78', 'Hi! I''m Diana-klm78 and I love traveling and gaming', 27, 'female', 'https://placeholder.com/user14.jpg'),
(15, 'Edward-nop90', 'Hi! I''m Edward-nop90 and I love reading and dancing', 36, 'male', 'https://placeholder.com/user15.jpg'),
(16, 'Fiona-qrs12', 'Hi! I''m Fiona-qrs12 and I love music and hiking', 30, 'female', 'https://placeholder.com/user16.jpg'),
(17, 'George-tuv34', 'Hi! I''m George-tuv34 and I love sports and photography', 32, 'male', 'https://placeholder.com/user17.jpg'),
(18, 'Hannah-wxy56', 'Hi! I''m Hannah-wxy56 and I love art and cooking', 28, 'female', 'https://placeholder.com/user18.jpg'),
(19, 'Ian-zab78', 'Hi! I''m Ian-zab78 and I love gaming and traveling', 31, 'male', 'https://placeholder.com/user19.jpg'),
(20, 'Julia-cde90', 'Hi! I''m Julia-cde90 and I love dancing and reading', 29, 'female', 'https://placeholder.com/user20.jpg'),
(21, 'Test-Profile', 'Hi! I''m Julia-cde90 and I love dancing and reading', 33, 'male', 'https://placeholder.com/user21.jpg');

INSERT INTO swipes (swiper_id, swiped_id, is_like) 
VALUES 
    (1, 21, true),
    (2, 21, true),
    (3, 21, true),
    (4, 21, true),
    (5, 21, true);