-- Drop indexes
DROP INDEX IF EXISTS idx_items_product_id;
DROP INDEX IF EXISTS idx_items_order_id;
DROP INDEX IF EXISTS idx_orders_status;
DROP INDEX IF EXISTS idx_orders_user_id;

-- Drop items table
DROP TABLE IF EXISTS items;

-- Drop orders table
DROP TABLE IF EXISTS orders;