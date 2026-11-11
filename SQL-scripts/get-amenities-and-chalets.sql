SELECT 
    c.title AS chalet_name,
    -- Собираем все названия удобств в один JSON-массив ['Wi-Fi', 'Баня']
    -- COALESCE нужен, чтобы вместо NULL (если у шале нет удобств) возвращался пустой массив []
    COALESCE(json_agg(a.title) FILTER (WHERE a.title IS NOT NULL), '[]') AS amenities
FROM chalets c
LEFT JOIN chalet_amenities ca ON c.id = ca.chalet_id
LEFT JOIN amenities a ON ca.amenity_id = a.id
GROUP BY c.id, c.title
ORDER BY chalet_name;
