-- phpMyAdmin SQL Dump
-- version 4.5.4.1deb2ubuntu2.1
-- http://www.phpmyadmin.net
--
-- Host: localhost
-- Generation Time: Jan 20, 2021 at 10:18 AM
-- Server version: 5.7.29
-- PHP Version: 7.0.33-0ubuntu0.16.04.7

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `gocommerce`
--

-- --------------------------------------------------------

--
-- Table structure for table `tbl_orders`
--

CREATE TABLE `tbl_orders` (
  `id` int(11) UNSIGNED NOT NULL,
  `product_id` int(11) DEFAULT NULL,
  `user_id` int(11) DEFAULT NULL,
  `quantity` int(11) DEFAULT NULL,
  `created_date` varchar(155) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `tbl_orders`
--

INSERT INTO `tbl_orders` (`id`, `product_id`, `user_id`, `quantity`, `created_date`) VALUES
(7, 2, 8, 1, 'Jan 18, 2021 12:12:37 PM'),
(8, 3, 8, 1, 'Jan 18, 2021 12:17:29 PM'),
(10, 2, 3, 1, 'Jan 18, 2021 01:16:39 PM'),
(11, 3, 3, 1, 'Jan 18, 2021 02:35:37 PM'),
(12, 2, 1, 1, 'Jan 18, 2021 02:36:01 PM'),
(13, 2, 2, 1, 'Jan 18, 2021 02:37:41 PM'),
(14, 1, 9, 1, 'Jan 18, 2021 02:54:07 PM'),
(15, 3, 2, 1, 'Jan 19, 2021 06:31:02 PM'),
(16, 1, 6, 1, 'Jan 20, 2021 10:08:51 AM'),
(17, 2, 6, 1, 'Jan 20, 2021 10:10:23 AM');

-- --------------------------------------------------------

--
-- Table structure for table `tbl_products`
--

CREATE TABLE `tbl_products` (
  `id` int(10) UNSIGNED NOT NULL,
  `title` varchar(255) DEFAULT NULL,
  `description` text,
  `price` float(10,2) DEFAULT NULL,
  `image` varchar(255) DEFAULT NULL,
  `image_path` varchar(255) DEFAULT NULL,
  `overall_rating` int(11) DEFAULT NULL,
  `wishlist_count` int(11) DEFAULT NULL,
  `created_date` datetime DEFAULT NULL,
  `modified_date` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `tbl_products`
--

INSERT INTO `tbl_products` (`id`, `title`, `description`, `price`, `image`, `image_path`, `overall_rating`, `wishlist_count`, `created_date`, `modified_date`) VALUES
(1, 'Bharat Lifestyle Fabric 3 + 1 + 1 Cream Sofa Set', 'Filling Material: Foam\r\nFrame Material: Solid Wood\r\ndiy- Basic assembly to be done with simple tools by the customer, comes with instructions.\r\nThis sofa set is a clear all-rounder', 10000.00, 'product-img1.jpg', 'assets/img/product-img1.jpg', NULL, 4, '2021-01-04 12:19:16', NULL),
(2, 'Hindhustan Lifestyle Fabric 3 + 1 + 1 Cream Sofa Set', 'Filling Material: cotton\r\nFrame Material: Oak Wood\r\ndiy- Basic assembly to be done with simple tools by the customer, comes with instructions.\r\nThis sofa set is a clear all-rounder', 8999.00, 'product-img2.jpg', 'assets/img/product-img2.jpg', NULL, 3, '2021-01-04 12:32:11', NULL),
(3, ' Chennai Lifestyle Fabric 3 + 2 + 1 Cream Sofa Set', 'Filling Material: Foam\r\nFrame Material: Solid Wood\r\nDIY - Basic assembly to be done with simple tools by the customer, comes with instructions.\r\nThis sofa set is a clear all-rounder', 11999.00, 'product-img3.jpg', 'assets/img/product-img3.jpg', NULL, 1, '2021-01-05 11:33:14', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `tbl_rating_reviews`
--

CREATE TABLE `tbl_rating_reviews` (
  `id` int(11) NOT NULL,
  `product_id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `rating` int(11) NOT NULL,
  `review` text,
  `created_date` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `tbl_user`
--

CREATE TABLE `tbl_user` (
  `id` int(10) UNSIGNED NOT NULL,
  `first_name` varchar(255) NOT NULL,
  `email_id` varchar(255) NOT NULL,
  `mobile_number` varchar(45) NOT NULL,
  `password` varchar(455) DEFAULT NULL,
  `created_date` varchar(155) DEFAULT NULL,
  `modified_date` varchar(155) DEFAULT NULL,
  `last_login_date` varchar(155) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `tbl_user`
--

INSERT INTO `tbl_user` (`id`, `first_name`, `email_id`, `mobile_number`, `password`, `created_date`, `modified_date`, `last_login_date`) VALUES
(1, 'Sakthivel', 'piccotvm10@gmail.com', '8870441514', '4b7eb3ed9a6bd5feb8ba4f076ec7e44acfa23332f49a0085e333adc45a0110346ca2f96b97a1', '2021-01-02 07:45:27', NULL, 'Jan 19, 2021 03:55:20 PM'),
(2, 'Piccosoft', 'piccotvm11@gmail.com', '898989898', '1d9044cc4be9c6621a7a9ef17710f400a7dd247eba50a7dad71a026751a2080cdd5c3f3a48', '2021-01-04 10:25:42 AM', NULL, 'Jan 19, 2021 06:14:19 PM'),
(3, 'Mani', 'piccotvm12@gmail.com', '8668042726', '9c4a6e23cba911143baee7eebef5eda78df5aa3b8196d2e74e8101b00abe369c1c41ffc8', '2021-01-05 10:58:46 AM', NULL, 'Jan 19, 2021 03:01:48 PM'),
(4, 'Raja', 'piccotvm13@gmail.com', '8668042727', '8031f8d1f798570e3bb513157b3381b8b1101604454b4ad7dc0ed00ea3264e0c1217630bef', '2021-01-09 10:2:35 AM', NULL, '2021-01-13 12:33:26 PM'),
(5, 'meganathan', 'meganathan@gmail.com', '8110074732', '12c82b2df8f2961efa932e04021219f919cd3a939e2f6b937298baf8ca61c224fccc6f34', '2021-01-09 3:19:7 PM', NULL, '2021-01-12 7:0:28 PM'),
(6, 'Seenu', 'seenu@gmail.com', '8989898982', '0f72f5e66c0bf1ca69531972b53b49dc5849076f87673ca43af9c093cb470816ca76f6677b', '2021-01-11 12:15:31 PM', NULL, 'Jan 20, 2021 10:07:33 AM'),
(7, 'stalin', 'stalin@gmail.com', '8110074732', 'a93ee2b69eaa87535325e2339fa0a88d2a0264037a869b5f8e1744ed3adf79dfeab6f309b519', '2021-01-13 11:3:28 AM', NULL, '2021-01-13 11:3:29 AM'),
(8, 'Sanjai', 'sanjai@gmail.com', '8998668042', '88866dc06721dd1cbb3738ec41e8e94ec1269a4ca3d4591c6334831b783108d1fa45effba827', '2021-01-18 11:30:46 AM', NULL, '2021-01-18 11:30:47 AM'),
(9, 'Gokul Nath', 'gokul@gmai.com', '8765432109', 'a8d1aef3c0f5a1b9705202c97ac236151825cf801a51c25ed9a001d01aa9aba17f2062c749', 'Jan 18, 2021 02:53:38 PM', NULL, 'Jan 18, 2021 02:53:39 PM');

-- --------------------------------------------------------

--
-- Table structure for table `tbl_wishlist`
--

CREATE TABLE `tbl_wishlist` (
  `id` int(10) UNSIGNED NOT NULL,
  `product_id` int(11) DEFAULT NULL,
  `user_id` int(11) DEFAULT NULL,
  `created_date` varchar(155) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `tbl_wishlist`
--

INSERT INTO `tbl_wishlist` (`id`, `product_id`, `user_id`, `created_date`) VALUES
(20, 1, 8, 'Jan 18, 2021 12:51:04 PM'),
(21, 2, 8, 'Jan 18, 2021 12:52:05 PM'),
(42, 1, 9, 'Jan 18, 2021 02:53:55 PM'),
(50, 3, 3, 'Jan 19, 2021 03:20:03 PM'),
(61, 2, 1, 'Jan 19, 2021 03:55:32 PM'),
(62, 1, 1, 'Jan 19, 2021 04:16:51 PM'),
(63, 1, 2, 'Jan 19, 2021 06:31:26 PM'),
(64, 2, 2, 'Jan 19, 2021 06:57:43 PM');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `tbl_orders`
--
ALTER TABLE `tbl_orders`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `tbl_products`
--
ALTER TABLE `tbl_products`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `tbl_rating_reviews`
--
ALTER TABLE `tbl_rating_reviews`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `tbl_user`
--
ALTER TABLE `tbl_user`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `tbl_wishlist`
--
ALTER TABLE `tbl_wishlist`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `tbl_orders`
--
ALTER TABLE `tbl_orders`
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=18;
--
-- AUTO_INCREMENT for table `tbl_products`
--
ALTER TABLE `tbl_products`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;
--
-- AUTO_INCREMENT for table `tbl_rating_reviews`
--
ALTER TABLE `tbl_rating_reviews`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `tbl_user`
--
ALTER TABLE `tbl_user`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=10;
--
-- AUTO_INCREMENT for table `tbl_wishlist`
--
ALTER TABLE `tbl_wishlist`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=68;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
