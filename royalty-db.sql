-- MariaDB dump 10.19  Distrib 10.4.21-MariaDB, for osx10.10 (x86_64)
--
-- Host: 127.0.0.1    Database: test-db
-- ------------------------------------------------------
-- Server version	10.4.21-MariaDB

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `transactions`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `transactions` (
  `id` int(11) NOT NULL,
  `invoice_pubid` varchar(50) DEFAULT NULL,
  `amount` float DEFAULT NULL,
  `charge_amount` float DEFAULT NULL,
  `user_pubid` varchar(255) DEFAULT NULL,
  `create_ts` timestamp NULL DEFAULT NULL,
  `update_ts` timestamp NULL DEFAULT NULL,
  `invoice_type` enum('TENANT','SHOP') DEFAULT NULL,
  `is_eligible_voucher` tinyint(1) DEFAULT 1,
  PRIMARY KEY (`id`),
  UNIQUE KEY `transactions_pk` (`invoice_pubid`)
);
/*!40101 SET character_set_client = @saved_cs_client */;
--
-- Table structure for table `users`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `users` (
  `pubid` varchar(100) NOT NULL,
  `email` varchar(100) DEFAULT NULL,
  `mdn` varchar(15) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `username` varchar(25) DEFAULT NULL,
  `full_name` varchar(255) DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  `date_of_birth` date DEFAULT NULL,
  `gender` char(1) DEFAULT NULL,
  `create_ts` timestamp NULL DEFAULT NULL,
  `create_by` varchar(50) DEFAULT NULL,
  `update_ts` timestamp NULL DEFAULT NULL,
  `update_by` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`pubid`),
  KEY `users_mdn_email_index` (`mdn`,`email`)
);
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `vouchers`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `vouchers` (
  `id` int(11) NOT NULL,
  `start_ts` timestamp NULL DEFAULT NULL,
  `end_ts` timestamp NULL DEFAULT NULL,
  `total_budget` float DEFAULT NULL,
  `voucher_amount` float DEFAULT NULL,
  `create_ts` timestamp NULL DEFAULT NULL,
  `update_ts` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
);
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `vouchers_detail`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `vouchers_detail` (
  `id` int(11) NOT NULL,
  `voucher_id` int(11) DEFAULT NULL,
  `code` varchar(10) DEFAULT NULL,
  `status` enum('AVAILABLE','GIVEN','REDEEMED') DEFAULT NULL,
  `user_pubid` varchar(255) DEFAULT NULL,
  `given_ts` timestamp NULL DEFAULT NULL,
  `redeem_ts` timestamp NULL DEFAULT NULL,
  `transaction_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `vouchers_detail_vouchers_id_fk` (`voucher_id`),
  CONSTRAINT `vouchers_detail_vouchers_id_fk` FOREIGN KEY (`voucher_id`) REFERENCES `vouchers` (`id`)
);
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-03-10 11:44:22
