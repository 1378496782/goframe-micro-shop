# 全链路压测-Locust脚本开发方案

## 1. Locust压测框架概述

Locust是一个开源的、分布式的用户负载测试工具，使用Python编写，支持HTTP和WebSocket协议的压测。本方案将基于Locust设计覆盖电商系统核心业务流程的压测脚本。

## 2. 脚本架构设计

### 2.1 目录结构

```
locust_tests/
├── config/
│   ├── __init__.py
│   ├── settings.py        # 全局配置
│   └── test_data.py       # 测试数据配置
├── core/
│   ├── __init__.py
│   ├── base_user.py       # 基础用户类
│   ├── http_client.py     # 自定义HTTP客户端
│   └── utils.py           # 工具函数
├── tasks/
│   ├── __init__.py
│   ├── user_tasks.py      # 用户相关任务
│   ├── goods_tasks.py     # 商品相关任务
│   ├── order_tasks.py     # 订单相关任务
│   └── payment_tasks.py   # 支付相关任务
├── flows/
│   ├── __init__.py
│   ├── shopping_flow.py   # 完整购物流程
│   ├── search_flow.py     # 搜索流程
│   └── user_center_flow.py # 用户中心流程
├── data_generator/
│   ├── __init__.py
│   ├── user_generator.py  # 用户数据生成
│   └── goods_generator.py # 商品数据生成
├── locustfile.py          # 主入口文件
├── requirements.txt       # 依赖管理
└── README.md              # 使用说明
```

### 2.2 核心组件设计

#### 2.2.1 基础用户类

```python
# core/base_user.py
from locust import HttpUser, task, between
from locust.contrib.fasthttp import FastHttpUser
import json
import logging
from core.utils import generate_pressure_test_header

class BaseUser(FastHttpUser):
    wait_time = between(1, 3)  # 任务间隔时间
    abstract = True  # 抽象类，不直接实例化
    
    def on_start(self):
        # 用户初始化，登录等操作
        self.login()
    
    def login(self):
        # 登录逻辑，获取token
        login_data = {
            "username": "test_user",
            "password": "CHANGE_ME_SECRET"
        }
        headers = generate_pressure_test_header()  # 添加压测标记
        
        response = self.client.post("/api/user/login", 
                                  json=login_data,
                                  headers=headers)
        
        if response.status_code == 200:
            self.token = response.json()["data"]["token"]
            logging.info(f"User logged in successfully: {self.token[:10]}...")
        else:
            logging.error(f"Login failed: {response.status_code}")
    
    def request_with_auth(self, method, url, **kwargs):
        # 带认证的请求方法
        headers = kwargs.pop("headers", {})
        headers.update({
            "Authorization": f"Bearer {self.token}"
        })
        headers.update(generate_pressure_test_header())  # 添加压测标记
        
        return self.client.request(method, url, headers=headers, **kwargs)
```

#### 2.2.2 工具函数

```python
# core/utils.py
import random
import string
import time

def generate_pressure_test_header():
    """生成压测标记头"""
    return {
        "X-Pressure-Test": "true",
        "X-Pressure-Test-ID": generate_unique_id()
    }

def generate_unique_id():
    """生成唯一ID"""
    timestamp = str(int(time.time() * 1000))
    random_str = ''.join(random.choices(string.ascii_letters + string.digits, k=8))
    return f"PT-{timestamp}-{random_str}"

def generate_phone_number():
    """生成测试手机号"""
    prefix = ['130', '131', '132', '133', '134', '135', '136', '137', '138', '139']
    return random.choice(prefix) + ''.join(random.choices(string.digits, k=8))

def generate_order_no():
    """生成订单号"""
    return f"ORD{int(time.time() * 1000)}{random.randint(100, 999)}"
```

## 3. 核心业务流程覆盖

### 3.1 用户流程

```python
# tasks/user_tasks.py
from locust import task, between
from core.base_user import BaseUser
import logging

class UserTasks(BaseUser):
    
    @task(1)
    def get_user_info(self):
        """获取用户信息"""
        response = self.request_with_auth("GET", "/api/user/info")
        
        if response.status_code == 200:
            logging.info("Get user info success")
        else:
            logging.error(f"Get user info failed: {response.status_code}")
    
    @task(1)
    def update_user_profile(self):
        """更新用户资料"""
        data = {
            "nickname": f"TestUser_{random.randint(1000, 9999)}",
            "gender": random.choice([0, 1]),
            "avatar": "http://example.com/avatar.jpg"
        }
        
        response = self.request_with_auth("PUT", "/api/user/profile", json=data)
        
        if response.status_code == 200:
            logging.info("Update user profile success")
        else:
            logging.error(f"Update user profile failed: {response.status_code}")
```

### 3.2 商品流程

```python
# tasks/goods_tasks.py
from locust import task, between
from core.base_user import BaseUser
import logging
import random

class GoodsTasks(BaseUser):
    
    @task(3)
    def get_goods_list(self):
        """获取商品列表"""
        params = {
            "page": 1,
            "size": 20,
            "category_id": random.randint(1, 10)
        }
        
        response = self.request_with_auth("GET", "/api/goods/list", params=params)
        
        if response.status_code == 200:
            logging.info("Get goods list success")
            data = response.json()
            if data.get("data", {}).get("list"):
                # 缓存商品ID，供后续使用
                self.goods_ids = [item["id"] for item in data["data"]["list"]]
        else:
            logging.error(f"Get goods list failed: {response.status_code}")
    
    @task(2)
    def get_goods_detail(self):
        """获取商品详情"""
        if hasattr(self, 'goods_ids') and self.goods_ids:
            goods_id = random.choice(self.goods_ids)
        else:
            goods_id = random.randint(100, 1000)
            
        response = self.request_with_auth("GET", f"/api/goods/detail/{goods_id}")
        
        if response.status_code == 200:
            logging.info(f"Get goods detail success: {goods_id}")
        else:
            logging.error(f"Get goods detail failed: {response.status_code}")
```

### 3.3 订单流程

```python
# tasks/order_tasks.py
from locust import task, between
from core.base_user import BaseUser
from core.utils import generate_order_no
import logging
import random

class OrderTasks(BaseUser):
    
    def on_start(self):
        super().on_start()
        # 初始化订单数据
        self.prepare_order_data()
    
    def prepare_order_data(self):
        """准备订单数据"""
        self.goods_items = [
            {"goods_id": random.randint(100, 1000), "quantity": 1, "price": random.randint(100, 10000)},
            {"goods_id": random.randint(100, 1000), "quantity": 2, "price": random.randint(100, 10000)}
        ]
        self.address_id = random.randint(1, 100)
    
    @task(1)
    def create_order(self):
        """创建订单"""
        order_data = {
            "number": generate_order_no(),
            "goods_items": self.goods_items,
            "address_id": self.address_id,
            "total_amount": sum(item["price"] * item["quantity"] for item in self.goods_items),
            "pay_type": 1  # 微信支付
        }
        
        response = self.request_with_auth("POST", "/api/order/create", json=order_data)
        
        if response.status_code == 200:
            order_id = response.json()["data"]["id"]
            logging.info(f"Create order success: {order_id}")
            self.last_order_id = order_id
        else:
            logging.error(f"Create order failed: {response.status_code}")
    
    @task(1)
    def get_order_list(self):
        """获取订单列表"""
        params = {
            "page": 1,
            "size": 10,
            "status": random.choice([0, 1, 2, 3, 4])
        }
        
        response = self.request_with_auth("GET", "/api/order/list", params=params)
        
        if response.status_code == 200:
            logging.info("Get order list success")
        else:
            logging.error(f"Get order list failed: {response.status_code}")
```

### 3.4 支付流程

```python
# tasks/payment_tasks.py
from locust import task, between
from core.base_user import BaseUser
import logging
import time

class PaymentTasks(BaseUser):
    
    @task(1)
    def create_payment(self):
        """创建支付"""
        # 假设已经有订单ID
        order_id = getattr(self, 'last_order_id', 1000)
        
        payment_data = {
            "order_id": order_id,
            "amount": 9900,  # 99元，单位分
            "pay_type": 1,
            "trade_no": f"PAY{int(time.time() * 1000)}"
        }
        
        response = self.request_with_auth("POST", "/api/payment/create", json=payment_data)
        
        if response.status_code == 200:
            logging.info(f"Create payment success for order: {order_id}")
        else:
            logging.error(f"Create payment failed: {response.status_code}")
    
    @task(0.5)
    def create_refund(self):
        """创建退款"""
        # 假设已经有支付的订单ID
        order_id = getattr(self, 'last_order_id', 1000)
        
        refund_data = {
            "order_id": order_id,
            "goods_id": 100,
            "reason": "商品质量问题",
            "refund_amount": 9900
        }
        
        response = self.request_with_auth("POST", "/api/refund/create", json=refund_data)
        
        if response.status_code == 200:
            logging.info(f"Create refund success for order: {order_id}")
        else:
            logging.error(f"Create refund failed: {response.status_code}")
```

### 3.5 完整购物流程

```python
# flows/shopping_flow.py
from locust import task, between
from core.base_user import BaseUser
from core.utils import generate_order_no
import logging
import random
import time

class ShoppingFlowUser(BaseUser):
    wait_time = between(2, 5)  # 购物流程间隔稍长
    
    @task
    def complete_shopping_flow(self):
        """完整购物流程：浏览商品 -> 加入购物车 -> 下单 -> 支付"""
        try:
            # 1. 获取商品列表
            self.browse_goods_list()
            
            # 2. 查看商品详情
            goods_id = self.view_goods_detail()
            
            # 3. 加入购物车
            self.add_to_cart(goods_id)
            
            # 4. 创建订单
            order_id = self.create_order(goods_id)
            
            # 5. 支付订单
            self.pay_order(order_id)
            
            logging.info(f"Complete shopping flow success: order {order_id}")
            
        except Exception as e:
            logging.error(f"Shopping flow failed: {str(e)}")
    
    def browse_goods_list(self):
        """浏览商品列表"""
        response = self.request_with_auth("GET", "/api/goods/list?page=1&size=20")
        response.raise_for_status()
    
    def view_goods_detail(self):
        """查看商品详情"""
        goods_id = random.randint(100, 1000)
        response = self.request_with_auth("GET", f"/api/goods/detail/{goods_id}")
        response.raise_for_status()
        return goods_id
    
    def add_to_cart(self, goods_id):
        """加入购物车"""
        data = {
            "goods_id": goods_id,
            "quantity": 1
        }
        response = self.request_with_auth("POST", "/api/cart/add", json=data)
        response.raise_for_status()
    
    def create_order(self, goods_id):
        """创建订单"""
        order_data = {
            "number": generate_order_no(),
            "goods_items": [{"goods_id": goods_id, "quantity": 1, "price": 9900}],
            "address_id": 1,
            "total_amount": 9900,
            "pay_type": 1
        }
        response = self.request_with_auth("POST", "/api/order/create", json=order_data)
        response.raise_for_status()
        return response.json()["data"]["id"]
    
    def pay_order(self, order_id):
        """支付订单"""
        payment_data = {
            "order_id": order_id,
            "amount": 9900,
            "pay_type": 1,
            "trade_no": f"PAY{int(time.time() * 1000)}"
        }
        response = self.request_with_auth("POST", "/api/payment/create", json=payment_data)
        response.raise_for_status()
```

## 4. 主入口文件

```python
# locustfile.py
from locust import HttpUser, task, between, TaskSet, tag
from flows.shopping_flow import ShoppingFlowUser
from tasks.user_tasks import UserTasks
from tasks.goods_tasks import GoodsTasks
from tasks.order_tasks import OrderTasks
from tasks.payment_tasks import PaymentTasks

def run_locust():
    # 主入口，可用于命令行执行
    import locust.main
    locust.main.main()

# 混合压测用户类
class MixedLoadUser(ShoppingFlowUser, UserTasks, GoodsTasks, OrderTasks, PaymentTasks):
    """混合负载用户，同时执行多种任务"""
    
    # 任务权重配置
    tasks = {
        ShoppingFlowUser.complete_shopping_flow: 1,
        UserTasks.get_user_info: 2,
        GoodsTasks.get_goods_list: 3,
        OrderTasks.get_order_list: 1,
        PaymentTasks.create_payment: 1
    }

# 专项压测用户类
class OrderLoadUser(OrderTasks):
    """订单流程专项压测"""
    pass

class PaymentLoadUser(PaymentTasks):
    """支付流程专项压测"""
    pass

if __name__ == "__main__":
    run_locust()
```

## 5. 测试数据准备

### 5.1 数据生成器

```python
# data_generator/user_generator.py
import json
import random
import string
import csv
import os

def generate_test_users(count=100):
    """生成测试用户数据"""
    users = []
    for i in range(count):
        user = {
            "username": f"test_user_{i:04d}",
            "password": "CHANGE_ME_SECRET",
            "nickname": f"测试用户{i:04d}",
            "phone": f"138{random.randint(10000000, 99999999)}",
            "email": f"test_user_{i:04d}@example.com"
        }
        users.append(user)
    return users

def save_users_to_csv(users, filename="test_users.csv"):
    """保存用户数据到CSV文件"""
    with open(filename, 'w', newline='', encoding='utf-8') as f:
        fieldnames = ['username', 'password', 'nickname', 'phone', 'email']
        writer = csv.DictWriter(f, fieldnames=fieldnames)
        writer.writeheader()
        writer.writerows(users)
    print(f"Generated {len(users)} test users and saved to {filename}")

def save_users_to_json(users, filename="test_users.json"):
    """保存用户数据到JSON文件"""
    with open(filename, 'w', encoding='utf-8') as f:
        json.dump(users, f, ensure_ascii=False, indent=2)
    print(f"Generated {len(users)} test users and saved to {filename}")

# 使用示例
if __name__ == "__main__":
    users = generate_test_users(1000)
    save_users_to_csv(users)
    save_users_to_json(users)
```

### 5.2 压测数据初始化

```python
# data_generator/init_test_data.py
import requests
import json
import logging
from user_generator import generate_test_users

# 配置日志
logging.basicConfig(level=logging.INFO, format='%(asctime)s - %(levelname)s - %(message)s')

def init_test_data(base_url, user_count=100):
    """初始化测试数据"""
    # 1. 生成并创建测试用户
    users = generate_test_users(user_count)
    created_users = []
    
    for user in users:
        try:
            # 添加压测标记
            headers = {"X-Pressure-Test": "true"}
            
            # 调用用户注册接口
            response = requests.post(
                f"{base_url}/api/user/register",
                json=user,
                headers=headers
            )
            
            if response.status_code == 200:
                created_users.append(user)
                logging.info(f"Created user: {user['username']}")
            else:
                logging.error(f"Failed to create user {user['username']}: {response.text}")
                
        except Exception as e:
            logging.error(f"Error creating user {user['username']}: {str(e)}")
    
    logging.info(f"Successfully created {len(created_users)} users out of {user_count}")
    return created_users

# 使用示例
if __name__ == "__main__":
    BASE_URL = "http://localhost:8000"
    users = init_test_data(BASE_URL, 1000)
    # 保存创建的用户数据
    with open("created_test_users.json", "w") as f:
        json.dump(users, f, indent=2)
```

## 6. 压测场景设计

### 6.1 常规场景

1. **基准测试**
   - 50个并发用户，持续5分钟
   - 目标：验证系统稳定性，建立性能基准

2. **峰值测试**
   - 从100用户逐步增加到500用户，持续10分钟
   - 目标：找到系统瓶颈和最大承载能力

3. **持久化测试**
   - 200个并发用户，持续30分钟
   - 目标：验证系统长时间运行的稳定性

### 6.2 核心流程场景

1. **购物流程压测**
   - 100个用户同时执行完整购物流程
   - 重点监控：订单创建、支付处理性能

2. **支付高峰场景**
   - 模拟促销活动，300个用户同时支付
   - 重点监控：支付成功率、响应时间

3. **退款流程场景**
   - 100个用户同时提交退款申请
   - 重点监控：退款处理性能、数据库负载

## 7. 执行命令示例

### 7.1 启动Locust Web界面

```bash
# 基本启动方式
locust -f locustfile.py

# 指定主机和用户类
locust -f locustfile.py --host=http://localhost:8000 --class-picker

# 分布式压测 - 主控
locust -f locustfile.py --master

# 分布式压测 - 工作节点
locust -f locustfile.py --worker --master-host=127.0.0.1
```

### 7.2 非UI模式执行

```bash
# 非UI模式，固定用户数
locust -f locustfile.py --host=http://localhost:8000 --headless -u 100 -r 10 -t 5m --csv=results

# 使用特定用户类
locust -f locustfile.py ShoppingFlowUser --headless -u 200 -r 20 -t 30m --csv=shopping_flow_results
```

## 8. 结果分析

### 8.1 关键指标

- **响应时间**：平均响应时间、95%/99%响应时间
- **吞吐量**：每秒请求数(RPS)
- **错误率**：请求失败百分比
- **并发用户数**：同时在线用户数
- **系统资源**：CPU、内存、网络IO、磁盘IO

### 8.2 结果解读

1. **响应时间分析**：
   - 平均响应时间应小于预期阈值（如200ms）
   - 95%响应时间不应超过2倍平均响应时间
   - 99%响应时间不应超过3倍平均响应时间

2. **吞吐量分析**：
   - 计算系统的最大吞吐量
   - 分析不同接口的吞吐量分布

3. **错误率分析**：
   - 错误率应低于0.1%
   - 分析错误类型和原因

此方案提供了完整的Locust压测脚本开发框架，覆盖了电商系统的核心业务流程，可以根据实际需求进行扩展和调整。