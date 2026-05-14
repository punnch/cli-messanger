CREATE TABLE messanger.messages (
    id          UUID           PRIMARY KEY,
    user_id     UUID           NOT NULL     REFERENCES messanger.users(id) ON DELETE CASCADE,
    room_id     UUID           NOT NULL     REFERENCES messanger.rooms(id) ON DELETE CASCADE,
    content     TEXT           NOT NULL,
    created_at  TIMESTAMPTZ    NOT NULL
);
