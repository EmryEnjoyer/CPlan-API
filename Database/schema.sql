CREATE TABLE IF NOT EXISTS account (
    id SERIAL PRIMARY KEY,
    name TEXT,
    date_created TIMESTAMP NOT NULL DEFAULT NOW(),
    active BOOLEAN NOT NULL DEFAULT TRUE
);

CREATE TABLE IF NOT EXISTS user_group (
    id SERIAL PRIMARY KEY,
    account_id INTEGER REFERENCES account(id),
    name TEXT,
    date_created TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS policy (
    id SERIAL PRIMARY KEY,
    action BIGINT NOT NULL DEFAULT 0,
    resource_table TEXT NOT NULL,
    resource_id INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS user_group_policy (
    user_group_id INTEGER NOT NULL REFERENCES user_group(id),
    policy_id INTEGER NOT NULL REFERENCES policy(id),
    PRIMARY KEY (user_group_id, policy_id)
);

CREATE TABLE IF NOT EXISTS cplan_user (
    id SERIAL PRIMARY KEY,
    account_id INTEGER NOT NULL REFERENCES account(id),
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    date_joined TIMESTAMP NOT NULL DEFAULT NOW(),
    active BOOLEAN NOT NULL DEFAULT TRUE
);

CREATE TABLE IF NOT EXISTS user_user_group (
    user_id INTEGER NOT NULL REFERENCES cplan_user(id),
    user_group_id INTEGER NOT NULL REFERENCES user_group(id),
    PRIMARY KEY (user_id, user_group_id)
);

CREATE TABLE IF NOT EXISTS time_budget (
    id SERIAL PRIMARY KEY,
    owner_id INTEGER NOT NULL REFERENCES cplan_user(id),
    name TEXT NOT NULL,
    date_created TIMESTAMP NOT NULL DEFAULT NOW(),
    active BOOLEAN NOT NULL DEFAULT TRUE
);

CREATE TABLE IF NOT EXISTS time_budget_entry (
    id SERIAL PRIMARY KEY,
    time_budget_id INTEGER NOT NULL REFERENCES time_budget(id),
    duration INTERVAL NOT NULL,
    resource_table TEXT NOT NULL,
    resource_id INTEGER NOT NULL,
    date_created TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS time_budget_goal (
    id SERIAL PRIMARY KEY,
    time_budget_id INTEGER NOT NULL REFERENCES time_budget(id),
    duration INTERVAL NOT NULL,
    resource_table TEXT NOT NULL,
    resource_id INTEGER NOT NULL,
    date_created TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS project (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    date_created TIMESTAMP NOT NULL DEFAULT NOW(),
    account_id INTEGER NOT NULL REFERENCES account(id),
    active BOOLEAN NOT NULL DEFAULT TRUE
);

CREATE TABLE IF NOT EXISTS board (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    date_created TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS board_project (
    board_id INTEGER NOT NULL REFERENCES board(id),
    project_id INTEGER NOT NULL REFERENCES project(id),
    active BOOLEAN NOT NULL DEFAULT TRUE,
    PRIMARY KEY (board_id, project_id)
);

CREATE TABLE IF NOT EXISTS list_node (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    date_created TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS board_list_node (
    board_id INTEGER NOT NULL REFERENCES board(id),
    list_node_id INTEGER NOT NULL REFERENCES list_node(id),
    list_node_order INTEGER NOT NULL,
    PRIMARY KEY (board_id, list_node_id)
);

CREATE TYPE CONN_TYPE AS ENUM ('PARENT', 'DEPENDS', 'REFERENCES');

CREATE TABLE IF NOT EXISTS list_node_graph (
    source_id INTEGER REFERENCES list_node(id),
    destination_id INTEGER REFERENCES list_node(id),
    connection_type CONN_TYPE NOT NULL,
    PRIMARY KEY (source_id, destination_id)
);

CREATE TABLE IF NOT EXISTS list_node_ordering (
    source_id INTEGER REFERENCES list_node(id),
    destination_id INTEGER REFERENCES list_node(id),
    list_node_order INTEGER NOT NULL,
    PRIMARY KEY(source_id, destination_id)
);

CREATE TYPE MEDIA_TYPE AS ENUM (
    'image',
    'gif',
    'video',
    'pdf',
    'unknown'
);

CREATE TABLE IF NOT EXISTS media (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    media_type MEDIA_TYPE NOT NULL,
    media_location TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS list_node_media (
    list_node_id INTEGER NOT NULL REFERENCES list_node(id),
    media_id INTEGER NOT NULL REFERENCES media(id),
    PRIMARY KEY (list_node_id, media_id)
);

CREATE TABLE IF NOT EXISTS comment (
    id SERIAL PRIMARY KEY,
    list_node_id INTEGER NOT NULL REFERENCES list_node(id),
    author_id INTEGER NOT NULL REFERENCES cplan_user(id),
    content TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS comment_media (
    comment_id INTEGER NOT NULL REFERENCES comment(id),
    media_id INTEGER NOT NULL REFERENCES media(id),
    PRIMARY KEY (comment_id, media_id)
);

CREATE TABLE IF NOT EXISTS tag (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    color TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS list_node_tag (
    list_node_id INTEGER NOT NULL REFERENCES list_node(id),
    tag_id INTEGER NOT NULL REFERENCES tag(id),
    PRIMARY KEY (list_node_id, tag_id)
);