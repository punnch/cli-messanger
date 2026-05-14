CREATE TABLE messanger.rooms (
    id         UUID                        PRIMARY KEY,
    name       VARCHAR(50) UNIQUE NOT NULL CHECK(char_length(name) BETWEEN 1 AND 50),
    created_at TIMESTAMPTZ        NOT NULL
);

CREATE TABLE messanger.room_members (
    user_id UUID  NOT NULL REFERENCES messanger.users(id) ON DELETE CASCADE,
    room_id UUID  NOT NULL REFERENCES messanger.rooms(id) ON DELETE CASCADE,
    PRIMARY KEY   (user_id, room_id)
);
