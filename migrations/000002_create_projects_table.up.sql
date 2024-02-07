CREATE TABLE IF NOT EXISTS projects (
    id SERIAL PRIMARY KEY,
    public_id uuid UNIQUE DEFAULT gen_random_uuid() NOT NULL,
    name VARCHAR(50) NOT NULL,
    description VARCHAR(255) NOT NULL,
    owner_id INT NOT NULL,
    CONSTRAINT fk_users FOREIGN KEY(owner_id) REFERENCES users(id)
);
