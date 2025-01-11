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
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(swiper_id, swiped_id),
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
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX idx_subscriptions_user ON subscriptions(user_id);
CREATE INDEX idx_subscriptions_status ON subscriptions(status);