-- 查看所有数据库
SHOW DATABASES;

-- 查看当前使用的数据库
SELECT DATABASE();

-- 查看 MySQL 版本
SELECT VERSION();
-- 创建测试数据库
CREATE DATABASE IF NOT EXISTS go_test_db;
USE go_test_db;

CREATE TABLE websites 
CREATE TABLE access_log 
    

CREATE TABLE IF NOT EXISTS students(
    id int AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    age int,
    grade VARCHAR(20)
);
SELECT * FROM students
DROP TABLE IF EXISTS students;


INSERT INTO students(name,age,grade)
VALUES('刘明',20,'六年级')

INSERT INTO students (name, age, grade) VALUES
('李四', 19, '二年级'),
('王五', 18, '一年级'),
('赵六', 21, '四年级'),
('孙七', 17, '二年级');

DELETE  FROM students
WHERE id = 1;
DELETE FROM students; 

TRUNCATE TABLE students;   //只清空表内容并把id初始为1递增


 accounts 表（包含字段 id 主键， balance 账户余额）和 
 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。

CREATE TABLE IF NOT EXISTS accounts(
    id INT PRIMARY KEY AUTO_INCREMENT,
    balance INT NOT NULL
);
CREATE TABLE IF NOT EXISTS transactions(
    id INT PRIMARY KEY AUTO_INCREMENT,
    from_account_id  INT,
    to_account_id INT,
    amount INT NOT NULL
);

-- 执行修改
ALTER TABLE transactions 
MODIFY COLUMN from_account_id VARCHAR(50);

ALTER TABLE transactions 
MODIFY COLUMN to_account_id VARCHAR(50);

SELECT* FROM accounts;
SELECT* FROM transactions;

INSERT INTO accounts (id, balance) VALUES
(1, 1000),
(2, 500),
(3, 2000);

INSERT INTO transactions (id, from_account_id, to_account_id, amount) VALUES
(1, 'A', 'B', 100),
(2, 'D', 'C', 50),
(3, 'C', 'A', 200);
TRUNCATE TABLE accounts   //清空表内容并id从1开始
-- 执行修改
ALTER TABLE accounts 
MODIFY COLUMN id VARCHAR(20);

INSERT INTO accounts (id, balance) VALUES
('A', 1000),
('B', 500),
('C', 2000);

employees 表，包含字段 id 、 name 、 department 、 salary 。
CREATE TABLE IF NOT EXISTS employees (
    id INT PRIMARY KEY AuTO_INCREMENT,
    name VARCHAR(50) NOT NULL,
    department VARCHAR(50) NOT NULL,
    salary INT NOT NULL
);
INSERT INTO employees (name, department, salary) VALUES
('张三', '技术部', 15000),
('李四', '技术部', 18000),
('王五', '市场部', 12000),
('赵六', '技术部', 22000),
('孙七', '人事部', 10000);
select * from employees
CREATE DATABASE IF NOT EXISTS go_test_db;
USE go_test_db;

CREATE TABLE IF NOT EXISTS books (
    id INT PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(100) NOT NULL,
    author VARCHAR(100) NOT NULL,
    price DOUBLE NOT NULL
);
INSERT INTO books (title, author, price) VALUES
('Go语言入门', '张三', 45),
('Go并发编程', '李四', 60),
('区块链基础', '王五', 88),
('SQL实战', '赵六', 52);
select * from books;

-- 先删除子表
DROP TABLE IF EXISTS comments;
DROP TABLE IF EXISTS posts;
-- 再删除父表
DROP TABLE IF EXISTS users;

select* from users;
select* from posts;
select* from comments;

SELECT 
    p.id,
    p.title, 
    p.content,
    COUNT(c.id) as comment_count
FROM posts p
LEFT JOIN comments c ON p.id = c.post_id  -- 左连接：包括没有评论的文章
GROUP BY p.id, p.title, p.content
ORDER BY comment_count DESC
LIMIT 1;