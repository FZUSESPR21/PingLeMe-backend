SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Records of classes
-- ----------------------------
INSERT INTO `classes` VALUES (1, '2021-05-12 22:14:50.000', '2021-05-12 22:14:53.000', NULL, '2018级软件工程1班');
INSERT INTO `classes` VALUES (2, '2021-05-12 22:15:33.000', '2021-05-12 22:15:35.000', NULL, '2018级软件工程2班');
INSERT INTO `classes` VALUES (3, '2021-05-12 22:15:58.000', '2021-05-12 22:16:00.000', NULL, '2019级软件工程1班');
INSERT INTO `classes` VALUES (4, '2021-05-12 22:17:45.000', '2021-05-12 22:17:48.000', NULL, '2019级软件工程2班');


-- ----------------------------
-- Records of evaluation_item_scores
-- ----------------------------
INSERT INTO `evaluation_item_scores` VALUES (13, '2021-05-18 17:22:02.330', '2021-05-18 17:47:52.910', NULL, 21, 1, 10, 11);
INSERT INTO `evaluation_item_scores` VALUES (14, '2021-05-18 17:22:02.330', '2021-05-18 17:47:52.905', NULL, 22, 1, 5, 11);
INSERT INTO `evaluation_item_scores` VALUES (15, '2021-05-18 17:22:02.330', '2021-05-18 17:47:52.908', NULL, 23, 1, 5, 11);


-- ----------------------------
-- Records of evaluation_table_items
-- ----------------------------
INSERT INTO `evaluation_table_items` VALUES (21, '2021-05-18 16:03:24.134', '2021-05-18 16:03:24.134', NULL, 11, 'Table Col 1 Row 1', 10, 1, 0);
INSERT INTO `evaluation_table_items` VALUES (22, '2021-05-18 16:03:24.134', '2021-05-18 16:03:24.134', NULL, 11, 'Table Col 1 Row 2', 5, 2, 1);
INSERT INTO `evaluation_table_items` VALUES (23, '2021-05-18 16:03:24.134', '2021-05-18 16:03:24.134', NULL, 11, 'Table Col 2 Row 2', 5, 2, 2);


-- ----------------------------
-- Records of evaluation_tables
-- ----------------------------
INSERT INTO `evaluation_tables` VALUES (11, '2021-05-18 16:03:24.134', '2021-05-18 16:03:24.134', NULL, '12344', 1, 1);


-- ----------------------------
-- Records of final_evaluation_table_scores
-- ----------------------------
INSERT INTO `final_evaluation_table_scores` VALUES (1, '2021-05-18 17:47:21.089', '2021-05-18 17:47:52.914', NULL, 11, 1, 10);


-- ----------------------------
-- Records of pairs
-- ----------------------------
INSERT INTO `pairs` VALUES (1, '2021-05-12 22:19:11.000', '2021-05-12 22:19:14.000', NULL, 221801101, 221801201);
INSERT INTO `pairs` VALUES (2, '2021-05-12 22:23:25.000', '2021-05-12 22:23:28.000', NULL, 221801102, 221801202);


-- ----------------------------
-- Records of permissions
-- ----------------------------
INSERT INTO `permissions` VALUES (1, '2021-05-12 22:25:49.000', NULL, NULL, 'add_student', 1);
INSERT INTO `permissions` VALUES (2, NULL, NULL, NULL, 'delete_student', 2);
INSERT INTO `permissions` VALUES (3, NULL, NULL, NULL, 'create_team', 3);
INSERT INTO `permissions` VALUES (4, NULL, NULL, NULL, 'add_teacher', 4);


-- ----------------------------
-- Records of role_permission
-- ----------------------------
INSERT INTO `role_permission` VALUES (1, 1);
INSERT INTO `role_permission` VALUES (2, 3);


-- ----------------------------
-- Records of roles
-- ----------------------------
INSERT INTO `roles` VALUES (1, NULL, NULL, NULL, 1, 'teacher');
INSERT INTO `roles` VALUES (2, NULL, NULL, NULL, 2, 'assistant');
INSERT INTO `roles` VALUES (3, NULL, NULL, NULL, 9, 'admin');
INSERT INTO `roles` VALUES (4, NULL, NULL, NULL, 3, 'teamleader');


-- ----------------------------
-- Records of student_class
-- ----------------------------
INSERT INTO `student_class` VALUES (1, 13);
INSERT INTO `student_class` VALUES (1, 14);
INSERT INTO `student_class` VALUES (1, 15);
INSERT INTO `student_class` VALUES (2, 16);
INSERT INTO `student_class` VALUES (2, 18);
INSERT INTO `student_class` VALUES (3, 19);


-- ----------------------------
-- Records of student_team
-- ----------------------------
INSERT INTO `student_team` VALUES (1, 13);
INSERT INTO `student_team` VALUES (2, 14);
INSERT INTO `student_team` VALUES (3, 15);
INSERT INTO `student_team` VALUES (4, 16);
INSERT INTO `student_team` VALUES (4, 18);


-- ----------------------------
-- Records of teacher_class
-- ----------------------------
INSERT INTO `teacher_class` VALUES (1, 11);
INSERT INTO `teacher_class` VALUES (3, 11);
INSERT INTO `teacher_class` VALUES (2, 12);


-- ----------------------------
-- Records of teams
-- ----------------------------
INSERT INTO `teams` VALUES (1, '2021-05-11 22:02:06.000', '2021-05-11 22:02:12.000', NULL, 1, '评了么', 1, 1);
INSERT INTO `teams` VALUES (2, '2021-05-12 14:59:39.000', '2021-05-12 14:59:42.000', NULL, 2, '字节乱动', 2, 1);
INSERT INTO `teams` VALUES (3, '2021-05-12 15:00:33.000', '2021-05-12 15:00:35.000', NULL, 3, '逐梦校友圈', 3, 2);
INSERT INTO `teams` VALUES (4, '2021-05-12 15:01:41.000', '2021-05-12 15:01:45.000', NULL, 4, '烤盐人', 4, 2);


-- ----------------------------
-- Records of user_role
-- ----------------------------
INSERT INTO `user_role` VALUES (2, 1);
INSERT INTO `user_role` VALUES (2, 6);
INSERT INTO `user_role` VALUES (2, 8);
INSERT INTO `user_role` VALUES (1, 11);
INSERT INTO `user_role` VALUES (1, 12);


-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (1, '2021-05-11 19:53:29.508', '2021-05-11 19:53:29.508', NULL, '920001', '$12$dw.uNa0xcmhTBXI/R8Q4geffWroVgh5BJkqVi5jgF10xZiIvXruZq', '蔡琳', 2);
INSERT INTO `users` VALUES (6, '2021-05-12 21:39:44.341', '2021-05-12 21:39:44.341', NULL, '920002', '$12$dw.uNa0xcmhTBXI/R8Q4geffWroVgh5BJkqVi5jgF10xZiIvXruZq', '林芬芬', 2);
INSERT INTO `users` VALUES (8, '2021-05-12 21:41:09.809', '2021-05-12 21:41:09.809', NULL, '920003', '$12$dw.uNa0xcmhTBXI/R8Q4geffWroVgh5BJkqVi5jgF10xZiIvXruZq', '江兰帆', 2);
INSERT INTO `users` VALUES (11, '2021-05-12 21:50:59.358', '2021-05-12 21:50:59.358', NULL, '910001', '$12$dw.uNa0xcmhTBXI/R8Q4geffWroVgh5BJkqVi5jgF10xZiIvXruZq', '江兰帆', 1);
INSERT INTO `users` VALUES (12, '2021-05-12 21:52:06.160', '2021-05-12 21:52:06.160', NULL, '910002', '$12$dw.uNa0xcmhTBXI/R8Q4geffWroVgh5BJkqVi5jgF10xZiIvXruZq', '王灿辉', 1);
INSERT INTO `users` VALUES (13, '2021-05-12 21:55:50.000', '2021-05-12 21:58:45.000', NULL, '221801101', '$12$dw.uNa0xcmhTBXI/R8Q4geffWroVgh5BJkqVi5jgF10xZiIvXruZq', '黄志军', 0);
INSERT INTO `users` VALUES (14, '2021-05-12 21:58:58.000', '2021-05-12 21:59:07.000', NULL, '221801102', '$12$dw.uNa0xcmhTBXI/R8Q4geffWroVgh5BJkqVi5jgF10xZiIvXruZq', '林苍凉', 0);
INSERT INTO `users` VALUES (15, '2021-05-12 22:01:49.000', '2021-05-12 22:01:51.000', NULL, '221801103', '$12$dw.uNa0xcmhTBXI/R8Q4geffWroVgh5BJkqVi5jgF10xZiIvXruZq', '钟叶', 0);
INSERT INTO `users` VALUES (16, '2021-05-12 22:02:48.000', '2021-05-12 22:02:52.000', NULL, '221801201', '$12$dw.uNa0xcmhTBXI/R8Q4geffWroVgh5BJkqVi5jgF10xZiIvXruZq', '郑母鸡', 0);
INSERT INTO `users` VALUES (17, '2021-05-12 22:04:03.000', '2021-05-12 22:04:06.000', NULL, '10001', '$12$dw.uNa0xcmhTBXI/R8Q4geffWroVgh5BJkqVi5jgF10xZiIvXruZq', '超管林某', 9);
INSERT INTO `users` VALUES (18, '2021-05-12 22:19:57.000', '2021-05-12 22:19:59.000', NULL, '221801202', '$12$dw.uNa0xcmhTBXI/R8Q4geffWroVgh5BJkqVi5jgF10xZiIvXruZq', '吴帅', 0);
INSERT INTO `users` VALUES (19, '2021-05-12 22:21:48.000', '2021-05-12 22:21:51.000', NULL, '221901101', '$12$dw.uNa0xcmhTBXI/R8Q4geffWroVgh5BJkqVi5jgF10xZiIvXruZq', '张振安', 0);


SET FOREIGN_KEY_CHECKS = 1;
