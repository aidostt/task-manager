CREATE TABLE Tasks
(
    id      UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES Users(id),
    title VARCHAR(255) NOT NULL,
    description VARCHAR(500),
    status varchar(50) CHECK (status IN ('todo', 'in_progress', 'done')) default 'todo',
    priority varchar(50) CHECK (priority IN ('low', 'medium', 'high')) default 'low',
    created_at TIMESTAMP DEFAULT now()
);