set session my.number_of_products = '30';

INSERT INTO product
select
    id,
    concat('Product ', id)
FROM
    GENERATE_SERIES(
        1,
        current_setting('my.number_of_products'):: int
    ) as id;