-- +goose Up
-- +goose StatementBegin
-- Table: tasks
CREATE TABLE tasks IF NOT EXIST(
    id VARCHAR(255) PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    created_by VARCHAR(100) NOT NULL,
    assignee VARCHAR(100) NOT NULL,
    assigned_leader VARCHAR(100),
    status ENUM('Submitted', 'Revision','Updated', 'Approved', 'In Progress', 'Completed') NOT NULL DEFAULT 'Submitted',
    progress INT DEFAULT 0 CHECK (progress >= 0 AND progress <= 100),
    progress_by VARCHAR(100),
    deadline DATETIME,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Table: task_history
CREATE TABLE task_history IF NOT EXIST(
    id INT AUTO_INCREMENT PRIMARY KEY,
    task_id VARCHAR(255) NOT NULL,
    action_by VARCHAR(100) NOT NULL,
    action ENUM('Submitted', 'Revision','Updated', 'Approved', 'In Progress', 'Completed') NOT NULL,
    note TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
