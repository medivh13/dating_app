CREATE TABLE public.user_premium (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES public.user(id),
    package_id INT REFERENCES public.premium_package(id),
    purchase_date DATE,
    expiration_date DATE,
    is_active BOOLEAN DEFAULT true
);
