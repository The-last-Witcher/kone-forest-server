SELECT 
    c.id, 
    c.title, 
    c.description, 
    c.address, 
    c.price, 
    c.rating, 
    c.guestsmax, 
    c.bedrooms, 
    c.bathrooms, 
    c.is_active,
    -- Используем jsonb_agg и COALESCE для железной надежности
    COALESCE(jsonb_agg(a.title) FILTER (WHERE a.title IS NOT NULL), '[]'::jsonb) AS amenities
FROM chalets c
LEFT JOIN chalet_amenities ca ON c.id = ca.chalet_id
LEFT JOIN amenities a ON ca.amenity_id = a.id
GROUP BY 
    c.id, c.title, c.description, c.address, 
    c.price, c.rating, c.guestsmax, c.bedrooms, 
    c.bathrooms, c.is_active
ORDER BY c.id;
