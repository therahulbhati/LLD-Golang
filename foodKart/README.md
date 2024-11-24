# Foodkart: Online Food Ordering Service

## Description
Foodkart is a platform for online food ordering, enabling users to order food from serviceable restaurants in their area. Restaurants deliver food items to users' locations based on defined service areas.

## Features
- **Specialized Menus**: Each restaurant specializes in a single dish.
- **Serviceable Areas**: Restaurants can serve multiple areas.
- **Single Restaurant Ordering**: Users can place orders from one restaurant at a time, with multiple item quantities allowed.
- **Ratings and Reviews**: Users can rate restaurants (1 to 5) with optional comments.
- **Restaurant Rating**: Calculated as the average rating from all users.

## Functional Requirements

1. **User Registration**
    - `register_user(user_details)`: Registers a new user with `name`, `gender`, `phoneNumber` (unique), and `pincode`.

2. **User Login**
    - `login_user(user_id)`: Sets the context to the logged-in user for all subsequent operations. Logs out any previous user.

3. **Restaurant Registration**
    - `register_restaurant(restaurant_name, serviceable_pincodes, food_item_name, food_item_price, initial_quantity)`: Registers a restaurant under the logged-in user, specifying service areas, dish, price, and initial quantity.

4. **Update Quantity**
    - `update_quantity(restaurant_name, quantity_to_add)`: Allows restaurant owners to add more items to their existing quantity.

5. **Rating and Review**
    - `rate_restaurant(restaurant_name, rating, comment)`: Allows users to rate a restaurant and optionally leave a comment.

6. **Serviceable Restaurant List**
    - `show_restaurant(sort_by)`: Displays a list of serviceable restaurants along with the food item name and price.
        - Sort by **rating** or **price** in descending order.
    - Note: A restaurant is serviceable if it can deliver to the user's pincode and has a non-zero food quantity.

7. **Order Placement**
    - `place_order(restaurant_name, quantity)`: Places an order from a specified restaurant for the given quantity.

## Bonus Features
1. **Order History**
    - Provides the ability to fetch the order history for a given user.

## Additional Guidelines
- **In-memory Data**: Use an in-memory store instead of a database or NoSQL.
- **No UI**: Focus on backend functionality without a user interface.
- **Driver Class**: Write a driver class to demonstrate and test functionality.
- **Code Execution**: Ensure code compiles, runs, and completes all expected tasks.
- **No Internet Access**: Limit references strictly to syntax (no external resources).

## Expectations
1. **Code Quality**:
    - Modular, object-oriented design.
    - Separation of concerns.
    - Extensibility for future features.
2. **Error Handling**:
    - Handle edge cases and fail gracefully.
3. **Readability**:
    - Use meaningful names for variables, functions, and classes.
    - Proper indentation and code formatting.

## Sample Test Case

```plaintext
register_user("Pralove", "M", "phoneNumber-1", "HSR")
register_user("Nitesh", "M", "phoneNumber-2", "BTM")
register_user("Vatsal", "M", "phoneNumber-3", "BTM")

login_user("phoneNumber-1")

register_restaurant("Food Court-1", "BTM/HSR", "NI Thali", 100, 5)
register_restaurant("Food Court-2", "BTM", "Burger", 120, 3)

login_user("phoneNumber-2")
register_restaurant("Food Court-3", "HSR", "SI Thali", 150, 1)

login_user("phoneNumber-3")
show_restaurant("price")
```
