insert into roles (id, name, description) values ('b966bb1c-9240-437a-baf2-a0ee5fd3c7ce', 'guest', 'Guest');
insert into roles (id, name, description) values ('7b190952-0801-4130-8084-957973632e73', 'staff', 'Staff');
insert into roles (id, name, description) values ('2145886e-4941-49e1-935c-134b2790c474', 'admin', 'Administrator');

insert into organizations (id, name, address) values ('25132737-479a-4684-8622-29a18365d436', 'Organization 1', 'Address 1');
insert into organizations (id, name, address) values ('47836126-5627-436e-a626-671027670873', 'Organization 2', 'Address 2');

insert into divisions (id, name, description, organization_id) values ('132737-479a-4684-8622-29a18365d436', 'Division 1', 'Description 1', '25132737-479a-4684-8622-29a18365d436');
insert into divisions (id, name, description, organization_id) values ('25132737-479a-4684-8622-29a18365d436', 'Division 2', 'Description 2', '47836126-5627-436e-a626-671027670873');

insert into users (id, name, email, phone_number, id_card_number, organization_id, division_id) values ('23037452-49d6-424e-9832-823496864270', 'User Admin', 'user1@example.com', '1234567890', '1234567890', '25132737-479a-4684-8622-29a18365d436', '132737-479a-4684-8622-29a18365d436');
insert into users (id, name, email, phone_number, id_card_number, organization_id, division_id) values ('56789012-3456-7890-1234-567890123456', 'User Staff', 'user2@example.com', '1234567890', '1234567890', '47836126-5627-436e-a626-671027670873', '25132737-479a-4684-8622-29a18365d436');

insert into user_roles (id, user_id, role_id) values ('5b5dc6f4-afcc-4fb2-8143-aef770a90355', '23037452-49d6-424e-9832-823496864270', '7b190952-0801-4130-8084-957973632e73');
insert into user_roles (id, user_id, role_id) values ('26442f28-3e4a-4e78-b316-d72773c5abe7', '56789012-3456-7890-1234-567890123456', '2145886e-4941-49e1-935c-134b2790c474');

insert into rooms (id, name, description, is_available) values ('6d24aff6-2627-4454-905b-0bded4918fd7', '111', 'Room A', true);
insert into rooms (id, name, description, is_available) values ('d9dcb947-4b91-45a1-8e76-ef4995abb235', '222', 'Room B', true);
insert into rooms (id, name, description, is_available) values ('2ea05c09-ad0e-4b63-82f3-85f434113705', '333', 'Room C', true);
insert into rooms (id, name, description, is_available) values ('3531c840-ed37-4f38-9599-fa49a3bdeadc', '444', 'Room D', true);
insert into rooms (id, name, description, is_available) values ('2fa95d79-6544-47e5-ae4c-3007ca59b202', '555', 'Room E', true);
insert into rooms (id, name, description, is_available) values ('8b6e6e7f-1f7a-4216-8005-b78c01c655a8', '666', 'Room F', true);