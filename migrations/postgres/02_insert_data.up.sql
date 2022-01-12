INSERT INTO "project"("id", "name", "domain") VALUES ('e04766bc-3228-4cd9-bd22-09e3fa27a6be', 'UPM', 'upm.udevs.io');

INSERT INTO "client_platform"("id", "project_id", "name", "subdomain") VALUES ('7d4a4c38-dd84-4902-b744-0488b80a4c01', 'e04766bc-3228-4cd9-bd22-09e3fa27a6be', 'UPM ADMIN PANEL', 'admin.upm.udevs.io');
INSERT INTO "client_platform"("id", "project_id", "name", "subdomain") VALUES ('7d4a4c38-dd84-4902-b744-0488b80a4c02', 'e04766bc-3228-4cd9-bd22-09e3fa27a6be', 'UPM DEVELOPER PANEL', 'dev.upm.udevs.io');
INSERT INTO "client_platform"("id", "project_id", "name", "subdomain") VALUES ('7d4a4c38-dd84-4902-b744-0488b80a4c03', 'e04766bc-3228-4cd9-bd22-09e3fa27a6be', 'UPM GUEST PANEL', 'guest.upm.udevs.io');

INSERT INTO "client_type"("id", "name", "confirm_by", "self_register", "self_recover") VALUES ('5a3818a9-90f0-44e9-a053-3be0ba1e2c01', 'ADMIN', NULL, FALSE, FALSE);
INSERT INTO "client_type"("id", "name", "confirm_by", "self_register", "self_recover") VALUES ('5a3818a9-90f0-44e9-a053-3be0ba1e2c02', 'HR', 'PHONE', FALSE, TRUE);
INSERT INTO "client_type"("id", "name", "confirm_by", "self_register", "self_recover") VALUES ('5a3818a9-90f0-44e9-a053-3be0ba1e2c03', 'CEO', 'PHONE', FALSE, TRUE);
INSERT INTO "client_type"("id", "name", "confirm_by", "self_register", "self_recover") VALUES ('5a3818a9-90f0-44e9-a053-3be0ba1e2c04', 'DEV', 'EMAIL', FALSE, TRUE);
INSERT INTO "client_type"("id", "name", "confirm_by", "self_register", "self_recover") VALUES ('5a3818a9-90f0-44e9-a053-3be0ba1e2c05', 'GUEST', 'EMAIL', TRUE, TRUE);

INSERT INTO "relation"("id", "client_type_id", "type", "name", "description") VALUES ('2d4a4c38-90f0-44e9-b744-7be0ba1e2c01', '5a3818a9-90f0-44e9-a053-3be0ba1e2c03', 'BRANCH', 'UDEVS', 'Main Office');
INSERT INTO "relation"("id", "client_type_id", "type", "name", "description") VALUES ('2d4a4c38-90f0-44e9-b744-7be0ba1e2c02', '5a3818a9-90f0-44e9-a053-3be0ba1e2c03', 'BRANCH', 'BILLZ', 'BILLZ Team');
INSERT INTO "relation"("id", "client_type_id", "type", "name", "description") VALUES ('2d4a4c38-90f0-44e9-b744-7be0ba1e2c03', '5a3818a9-90f0-44e9-a053-3be0ba1e2c03', 'BRANCH', 'VENDOO', 'VENDOO Team');

INSERT INTO "user_info_field"("id", "client_type_id", "field_name", "field_type", "data_type") VALUES ('3a3818a9-90f0-44e9-b744-5be0ba1e2c01', '5a3818a9-90f0-44e9-a053-3be0ba1e2c04', 'resume_url', 'FLAT', 'STRING');
INSERT INTO "user_info_field"("id", "client_type_id", "field_name", "field_type", "data_type") VALUES ('3a3818a9-90f0-44e9-b744-5be0ba1e2c02', '5a3818a9-90f0-44e9-a053-3be0ba1e2c04', 'contanc_links', 'ARRAY', 'STRING');

INSERT INTO "client"("client_platform_id", "client_type_id", "login_strategy") VALUES ('7d4a4c38-dd84-4902-b744-0488b80a4c01', '5a3818a9-90f0-44e9-a053-3be0ba1e2c01', 'STANDARD');
INSERT INTO "client"("client_platform_id", "client_type_id", "login_strategy") VALUES ('7d4a4c38-dd84-4902-b744-0488b80a4c01', '5a3818a9-90f0-44e9-a053-3be0ba1e2c02', 'PASSCODE');
INSERT INTO "client"("client_platform_id", "client_type_id", "login_strategy") VALUES ('7d4a4c38-dd84-4902-b744-0488b80a4c01', '5a3818a9-90f0-44e9-a053-3be0ba1e2c03', 'ONE2MANY');
INSERT INTO "client"("client_platform_id", "client_type_id", "login_strategy") VALUES ('7d4a4c38-dd84-4902-b744-0488b80a4c02', '5a3818a9-90f0-44e9-a053-3be0ba1e2c04', 'STANDARD');
INSERT INTO "client"("client_platform_id", "client_type_id", "login_strategy") VALUES ('7d4a4c38-dd84-4902-b744-0488b80a4c03', '5a3818a9-90f0-44e9-a053-3be0ba1e2c05', 'OTP');

INSERT INTO "role"("id", "client_type_id", "name") VALUES ('a1ca1301-4da9-424d-a9e2-578ae6dcde01', '5a3818a9-90f0-44e9-a053-3be0ba1e2c01', 'DEFAULT');
INSERT INTO "role"("id", "client_type_id", "name") VALUES ('a1ca1301-4da9-424d-a9e2-578ae6dcde02', '5a3818a9-90f0-44e9-a053-3be0ba1e2c02', 'DEFAULT');
INSERT INTO "role"("id", "client_type_id", "name") VALUES ('a1ca1301-4da9-424d-a9e2-578ae6dcde03', '5a3818a9-90f0-44e9-a053-3be0ba1e2c03', 'DEFAULT');
INSERT INTO "role"("id", "client_type_id", "name") VALUES ('a1ca1301-4da9-424d-a9e2-578ae6dcde04', '5a3818a9-90f0-44e9-a053-3be0ba1e2c04', 'DEFAULT');
INSERT INTO "role"("id", "client_type_id", "name") VALUES ('a1ca1301-4da9-424d-a9e2-578ae6dcde05', '5a3818a9-90f0-44e9-a053-3be0ba1e2c05', 'DEFAULT');

INSERT INTO "role"("id", "client_type_id", "name") VALUES ('a1ca1301-4da9-424d-a9e2-578ae6dcde06', '5a3818a9-90f0-44e9-a053-3be0ba1e2c04', 'BackEnd');
INSERT INTO "role"("id", "client_type_id", "name") VALUES ('a1ca1301-4da9-424d-a9e2-578ae6dcde07', '5a3818a9-90f0-44e9-a053-3be0ba1e2c04', 'FrontEnd');
INSERT INTO "role"("id", "client_type_id", "name") VALUES ('a1ca1301-4da9-424d-a9e2-578ae6dcde08', '5a3818a9-90f0-44e9-a053-3be0ba1e2c04', 'QA');
INSERT INTO "role"("id", "client_type_id", "name") VALUES ('a1ca1301-4da9-424d-a9e2-578ae6dcde09', '5a3818a9-90f0-44e9-a053-3be0ba1e2c04', 'DevOps');
INSERT INTO "role"("id", "client_type_id", "name") VALUES ('a1ca1301-4da9-424d-a9e2-578ae6dcde10', '5a3818a9-90f0-44e9-a053-3be0ba1e2c04', 'Lead');
INSERT INTO "role"("id", "client_type_id", "name") VALUES ('a1ca1301-4da9-424d-a9e2-578ae6dcde11', '5a3818a9-90f0-44e9-a053-3be0ba1e2c04', 'PM');

INSERT INTO "role"("id", "client_type_id", "name") VALUES ('a1ca1301-4da9-424d-a9e2-578ae6dcde12', '5a3818a9-90f0-44e9-a053-3be0ba1e2c05', 'PO');
INSERT INTO "role"("id", "client_type_id", "name") VALUES ('a1ca1301-4da9-424d-a9e2-578ae6dcde13', '5a3818a9-90f0-44e9-a053-3be0ba1e2c05', 'Investor');

INSERT INTO "scope"("client_platform_id", "path", "method", "requests") VALUES ('7d4a4c38-dd84-4902-b744-0488b80a4c01', '/', 'GET', 123);
INSERT INTO "scope"("client_platform_id", "path", "method", "requests") VALUES ('7d4a4c38-dd84-4902-b744-0488b80a4c01', '/ping', 'GET', 777);
INSERT INTO "scope"("client_platform_id", "path", "method", "requests") VALUES ('7d4a4c38-dd84-4902-b744-0488b80a4c01', '/v1/user', 'POST', 11);
INSERT INTO "scope"("client_platform_id", "path", "method", "requests") VALUES ('7d4a4c38-dd84-4902-b744-0488b80a4c01', '/v1/user', 'GET', 22);
INSERT INTO "scope"("client_platform_id", "path", "method", "requests") VALUES ('7d4a4c38-dd84-4902-b744-0488b80a4c01', '/v1/user/:id', 'GET', 33);
INSERT INTO "scope"("client_platform_id", "path", "method", "requests") VALUES ('7d4a4c38-dd84-4902-b744-0488b80a4c01', '/v1/user/:id', 'PUT', 44);
INSERT INTO "scope"("client_platform_id", "path", "method", "requests") VALUES ('7d4a4c38-dd84-4902-b744-0488b80a4c01', '/v1/user/:id', 'DELETE', 55);

INSERT INTO "permission"("id", "client_platform_id", "parent_id", "name") VALUES ('ffffffff-ffff-4fff-8fff-ffffffffffff', '7d4a4c38-dd84-4902-b744-0488b80a4c01', NULL, '/root');
INSERT INTO "permission"("id", "client_platform_id", "parent_id", "name") VALUES ('9cbb32da-e473-4312-8413-95524ec08c01', '7d4a4c38-dd84-4902-b744-0488b80a4c01', 'ffffffff-ffff-4fff-8fff-ffffffffffff', '/root/dashboard');
INSERT INTO "permission"("id", "client_platform_id", "parent_id", "name") VALUES ('9cbb32da-e473-4312-8413-95524ec08c02', '7d4a4c38-dd84-4902-b744-0488b80a4c01', 'ffffffff-ffff-4fff-8fff-ffffffffffff', '/root/user');
INSERT INTO "permission"("id", "client_platform_id", "parent_id", "name") VALUES ('9cbb32da-e473-4312-8413-95524ec08c03', '7d4a4c38-dd84-4902-b744-0488b80a4c01', '9cbb32da-e473-4312-8413-95524ec08c02', '/root/user/list');
INSERT INTO "permission"("id", "client_platform_id", "parent_id", "name") VALUES ('9cbb32da-e473-4312-8413-95524ec08c04', '7d4a4c38-dd84-4902-b744-0488b80a4c01', '9cbb32da-e473-4312-8413-95524ec08c02', '/root/user/detail');
INSERT INTO "permission"("id", "client_platform_id", "parent_id", "name") VALUES ('9cbb32da-e473-4312-8413-95524ec08c05', '7d4a4c38-dd84-4902-b744-0488b80a4c01', '9cbb32da-e473-4312-8413-95524ec08c02', '/root/user/create');
INSERT INTO "permission"("id", "client_platform_id", "parent_id", "name") VALUES ('9cbb32da-e473-4312-8413-95524ec08c06', '7d4a4c38-dd84-4902-b744-0488b80a4c01', '9cbb32da-e473-4312-8413-95524ec08c02', '/root/user/update');
INSERT INTO "permission"("id", "client_platform_id", "parent_id", "name") VALUES ('9cbb32da-e473-4312-8413-95524ec08c07', '7d4a4c38-dd84-4902-b744-0488b80a4c01', '9cbb32da-e473-4312-8413-95524ec08c02', '/root/user/delete');

INSERT INTO "permission_scope"("permission_id", "client_platform_id", "path", "method") VALUES ('9cbb32da-e473-4312-8413-95524ec08c03', '7d4a4c38-dd84-4902-b744-0488b80a4c01', '/v1/user', 'GET');
INSERT INTO "permission_scope"("permission_id", "client_platform_id", "path", "method") VALUES ('9cbb32da-e473-4312-8413-95524ec08c05', '7d4a4c38-dd84-4902-b744-0488b80a4c01', '/v1/user', 'POST');
INSERT INTO "permission_scope"("permission_id", "client_platform_id", "path", "method") VALUES ('9cbb32da-e473-4312-8413-95524ec08c06', '7d4a4c38-dd84-4902-b744-0488b80a4c01', '/v1/user/:id', 'GET');
INSERT INTO "permission_scope"("permission_id", "client_platform_id", "path", "method") VALUES ('9cbb32da-e473-4312-8413-95524ec08c06', '7d4a4c38-dd84-4902-b744-0488b80a4c01', '/v1/user/:id', 'PUT');

INSERT INTO "role_permission"("role_id", "permission_id") VALUES ('a1ca1301-4da9-424d-a9e2-578ae6dcde01', 'ffffffff-ffff-4fff-8fff-ffffffffffff');
INSERT INTO "role_permission"("role_id", "permission_id") VALUES ('a1ca1301-4da9-424d-a9e2-578ae6dcde01', '9cbb32da-e473-4312-8413-95524ec08c01');
INSERT INTO "role_permission"("role_id", "permission_id") VALUES ('a1ca1301-4da9-424d-a9e2-578ae6dcde01', '9cbb32da-e473-4312-8413-95524ec08c02');
INSERT INTO "role_permission"("role_id", "permission_id") VALUES ('a1ca1301-4da9-424d-a9e2-578ae6dcde01', '9cbb32da-e473-4312-8413-95524ec08c03');
INSERT INTO "role_permission"("role_id", "permission_id") VALUES ('a1ca1301-4da9-424d-a9e2-578ae6dcde01', '9cbb32da-e473-4312-8413-95524ec08c04');
INSERT INTO "role_permission"("role_id", "permission_id") VALUES ('a1ca1301-4da9-424d-a9e2-578ae6dcde01', '9cbb32da-e473-4312-8413-95524ec08c05');
INSERT INTO "role_permission"("role_id", "permission_id") VALUES ('a1ca1301-4da9-424d-a9e2-578ae6dcde01', '9cbb32da-e473-4312-8413-95524ec08c06');
INSERT INTO "role_permission"("role_id", "permission_id") VALUES ('a1ca1301-4da9-424d-a9e2-578ae6dcde01', '9cbb32da-e473-4312-8413-95524ec08c07');

INSERT INTO "user"("id", "project_id", "client_platform_id", "client_type_id", "role_id", "phone", "email", "login", "password", "active", "expires_at", "created_at", "updated_at")
    VALUES ('f799f1c5-ce5f-4fdd-ac23-f542247dcc01', 'e04766bc-3228-4cd9-bd22-09e3fa27a6be', '7d4a4c38-dd84-4902-b744-0488b80a4c01', '5a3818a9-90f0-44e9-a053-3be0ba1e2c01', 'a1ca1301-4da9-424d-a9e2-578ae6dcde01', '+998943614198', 'saidamir.botirov@udevs.io', 'UPM!AD@22#00001', '$argon2id$v=19$models=65536,t=3,p=4$ylXXHk9qJ4Cg4te0mAjOAQ$sCU9Ie/L4SRHbd3sOKs+qwhImzlKZr4ya6waEaCO6Uw', 1, '2072-05-01T11:21:59.001+0000', '2022-01-17T11:21:59.001+0000', NOW());
INSERT INTO "user"("id", "project_id", "client_platform_id", "client_type_id", "role_id", "phone", "email", "login", "password", "active", "expires_at", "created_at", "updated_at")
    VALUES ('f799f1c5-ce5f-4fdd-ac23-f542247dcc02', 'e04766bc-3228-4cd9-bd22-09e3fa27a6be', '7d4a4c38-dd84-4902-b744-0488b80a4c02', '5a3818a9-90f0-44e9-a053-3be0ba1e2c04', 'a1ca1301-4da9-424d-a9e2-578ae6dcde02', '+998998465798', 'saidamir.botirov@gmail.com', 'UPM!DEV@22#00001', '$argon2id$v=19$models=65536,t=3,p=4$ylXXHk9qJ4Cg4te0mAjOAQ$sCU9Ie/L4SRHbd3sOKs+qwhImzlKZr4ya6waEaCO6Uw', 1, '2072-05-01T11:21:59.001+0000', '2022-01-17T11:21:59.001+0000', NOW());
INSERT INTO "user"("id", "project_id", "client_platform_id", "client_type_id", "role_id", "phone", "email", "login", "password", "active", "expires_at", "created_at", "updated_at")
    VALUES ('f799f1c5-ce5f-4fdd-ac23-f542247dcc03', 'e04766bc-3228-4cd9-bd22-09e3fa27a6be', '7d4a4c38-dd84-4902-b744-0488b80a4c03', '5a3818a9-90f0-44e9-a053-3be0ba1e2c05', 'a1ca1301-4da9-424d-a9e2-578ae6dcde03', '+998901234567', 'saidamir.botirov@gmail.com', 'UPM!GT@22#00001', '$argon2id$v=19$models=65536,t=3,p=4$ylXXHk9qJ4Cg4te0mAjOAQ$sCU9Ie/L4SRHbd3sOKs+qwhImzlKZr4ya6waEaCO6Uw', 1, '2072-05-01T11:21:59.001+0000', '2022-01-17T11:21:59.001+0000', NOW());