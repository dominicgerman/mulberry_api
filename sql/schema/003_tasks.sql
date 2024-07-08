-- +goose Up
CREATE TABLE tasks (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id),
    title VARCHAR(255) NOT NULL,
    notes TEXT, 
    frequency VARCHAR(50) NOT NULL, -- eg. 'daily', 'weekly', 'monthly', 'yearly' 
    next_due_date DATE NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE tasks;