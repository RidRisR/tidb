/*!40014 SET FOREIGN_KEY_CHECKS=0*/;
/*!40101 SET NAMES binary*/;
CREATE TABLE `pt_case_1` (
  `a` int DEFAULT NULL,
  `b` int DEFAULT NULL,
  UNIQUE KEY `idx` (`a`) /*T![global_index] GLOBAL */,
  KEY `idx1` (`a`) /*T![global_index] GLOBAL */
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin
PARTITION BY LIST (`b`)
(PARTITION `p0` VALUES IN (0,1,2,3),
 PARTITION `p1` VALUES IN (4,5,6),
 PARTITION `p2` VALUES IN (7,8,9,10));
