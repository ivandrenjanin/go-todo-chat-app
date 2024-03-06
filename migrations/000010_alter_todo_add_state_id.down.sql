ALTER TABLE
    todos DROP COLUMN state_id;

ALTER TABLE
    todos DROP COLUMN item_order;

ALTER TABLE
    todos DROP CONSTRAINT fk_todo_state;
