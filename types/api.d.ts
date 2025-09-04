// API响应基础类型
interface BaseResponse<T = any> {
  code: number;
  message: string;
  data: T;
}

// 分页响应类型
interface PaginatedResponse<T> {
  list: T[];
  total: number;
  page: number;
  size: number;
}

// 商品相关类型
interface Product {
  id: number;
  name: string;
  price: string;
  originalPrice: string;
  image: string;
  sales: number;
  stock?: number;
  description?: string;
}

interface ProductDetail extends Product {
  images: string[];
  specifications: Array<{
    name: string;
    values: string[];
  }>;
  description: string;
  stock: number;
}

// 分类类型
interface Category {
  id: number;
  name: string;
  icon: string;
}

// 用户相关类型
interface UserInfo {
  id: number;
  username: string;
  avatar: string;
  email?: string;
  phone?: string;
}

interface LoginRequest {
  username: string;
  password: string;
}

interface LoginResponse {
  token: string;
  userInfo: UserInfo;
}

// 购物车相关类型
interface CartItem {
  id: number;
  productId: number;
  productName: string;
  price: string;
  image: string;
  quantity: number;
  selected: boolean;
}

interface CartResponse {
  items: CartItem[];
  totalPrice: string;
  totalQuantity: number;
}

// 订单相关类型
interface OrderProduct {
  name: string;
  image: string;
  price: string;
  quantity: number;
}

interface Order {
  id: string;
  status: number; // 1: 待付款, 2: 待发货, 3: 已发货, 4: 已完成, 5: 已取消
  totalAmount: string;
  createTime: string;
  products: OrderProduct[];
}

// API方法类型定义
declare namespace API {
  // 商品相关
  function getGoodsList(params: { page: number; size: number }): Promise<BaseResponse<PaginatedResponse<Product>>>;
  function getGoodsDetail(id: number): Promise<BaseResponse<ProductDetail>>;
  
  // 分类相关
  function getCategories(): Promise<BaseResponse<Category[]>>;
  
  // 用户相关
  function login(data: LoginRequest): Promise<BaseResponse<LoginResponse>>;
  function getUserInfo(): Promise<BaseResponse<UserInfo>>;
  function register(data: any): Promise<BaseResponse>;
  function updatePassword(data: any): Promise<BaseResponse>;
  
  // 购物车相关
  function getCart(): Promise<BaseResponse<CartResponse>>;
  function addToCart(data: any): Promise<BaseResponse>;
  function updateCartItem(data: any): Promise<BaseResponse>;
  function removeCartItem(id: number): Promise<BaseResponse>;
  
  // 订单相关
  function createOrder(data: any): Promise<BaseResponse>;
  function getOrders(params: any): Promise<BaseResponse<PaginatedResponse<Order>>>;
  function getOrderDetail(id: string): Promise<BaseResponse<Order>>;
  function cancelOrder(id: string): Promise<BaseResponse>;
}