create table users (
    id int unsigned auto_increment primary key ,
    name varchar(255),

    updated_at       timestamp    null,
    created_at       timestamp    null
);

create table chats (
    id int unsigned auto_increment primary key,
    initiator_id int unsigned,
    second_id int unsigned,

    updated_at       timestamp    null,
    created_at       timestamp    null,

    constraint receiver_chat_id_fk foreign key (initiator_id) references users (id),
    constraint sender_chat_id_fk foreign key (second_id) references users (id)
);

create table messages (
    id int unsigned auto_increment primary key ,
    chat_id int unsigned,
    sender_id int unsigned,

    message varchar(1000),

    updated_at       timestamp    null,
    created_at       timestamp    null,

    constraint sender_id_message_fk foreign key (sender_id) references users (id),
    constraint chat_id_message_fk foreign key (chat_id) references chats (id)
)