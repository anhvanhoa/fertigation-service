-- Seed data for fertilizer_schedules table
INSERT INTO fertilizer_schedules (
    id, planting_cycle_id, fertilizer_type_id, application_date, dosage, 
    unit, application_method, growth_stage, weather_conditions, soil_conditions, 
    is_completed, completed_date, actual_dosage, effectiveness_rating, notes, created_by
) VALUES 
-- Seedling stage applications
('770e8400-e29b-41d4-a716-446655440001', '880e8400-e29b-41d4-a716-446655440001', '550e8400-e29b-41d4-a716-446655440001', 
 '2024-01-15', 0.5, 'kg', 'soil', 'seedling', 'sunny, 25°C', 'loamy, pH 6.5', 
 true, '2024-01-15', 0.5, 4, 'Applied compost for root development', '550e8400-e29b-41d4-a716-446655440000'),

('770e8400-e29b-41d4-a716-446655440002', '880e8400-e29b-41d4-a716-446655440002', '550e8400-e29b-41d4-a716-446655440002', 
 '2024-01-20', 0.3, 'kg', 'soil', 'seedling', 'cloudy, 22°C', 'clay, pH 6.8', 
 true, '2024-01-20', 0.3, 3, 'Chicken manure applied after composting', '550e8400-e29b-41d4-a716-446655440000'),

-- Vegetative stage applications
('770e8400-e29b-41d4-a716-446655440003', '880e8400-e29b-41d4-a716-446655440003', '550e8400-e29b-41d4-a716-446655440003', 
 '2024-02-01', 0.1, 'g', 'foliar', 'vegetative', 'sunny, 28°C', 'sandy loam, pH 6.2', 
 true, '2024-02-01', 0.1, 5, 'Excellent growth response to NPK foliar spray', '550e8400-e29b-41d4-a716-446655440000'),

('770e8400-e29b-41d4-a716-446655440004', '880e8400-e29b-41d4-a716-446655440004', '550e8400-e29b-41d4-a716-446655440004', 
 '2024-02-10', 0.05, 'g', 'soil', 'vegetative', 'partly cloudy, 26°C', 'loamy, pH 6.5', 
 true, '2024-02-10', 0.05, 4, 'Urea applied for nitrogen boost', '550e8400-e29b-41d4-a716-446655440000'),

-- Flowering stage applications
('770e8400-e29b-41d4-a716-446655440005', '880e8400-e29b-41d4-a716-446655440005', '550e8400-e29b-41d4-a716-446655440005', 
 '2024-03-01', 0.02, 'ml', 'foliar', 'flowering', 'sunny, 30°C', 'sandy, pH 6.0', 
 true, '2024-03-01', 0.02, 5, 'Seaweed extract enhanced flowering', '550e8400-e29b-41d4-a716-446655440000'),

('770e8400-e29b-41d4-a716-446655440006', '880e8400-e29b-41d4-a716-446655440006', '550e8400-e29b-41d4-a716-446655440006', 
 '2024-03-05', 0.1, 'ml', 'soil', 'flowering', 'cloudy, 27°C', 'clay loam, pH 6.3', 
 true, '2024-03-05', 0.1, 4, 'Fish emulsion improved soil biology', '550e8400-e29b-41d4-a716-446655440000'),

-- Fruiting stage applications
('770e8400-e29b-41d4-a716-446655440007', '880e8400-e29b-41d4-a716-446655440007', '550e8400-e29b-41d4-a716-446655440007', 
 '2024-03-15', 0.2, 'g', 'soil', 'fruiting', 'sunny, 32°C', 'loamy, pH 6.4', 
 true, '2024-03-15', 0.2, 5, 'Superphosphate enhanced fruit development', '550e8400-e29b-41d4-a716-446655440000'),

('770e8400-e29b-41d4-a716-446655440008', '880e8400-e29b-41d4-a716-446655440008', '550e8400-e29b-41d4-a716-446655440008', 
 '2024-03-20', 0.1, 'g', 'soil', 'fruiting', 'partly cloudy, 29°C', 'sandy loam, pH 6.1', 
 true, '2024-03-20', 0.1, 4, 'Potassium sulfate improved fruit quality', '550e8400-e29b-41d4-a716-446655440000'),

-- Hydroponic applications
('770e8400-e29b-41d4-a716-446655440009', '880e8400-e29b-41d4-a716-446655440009', '550e8400-e29b-41d4-a716-446655440009', 
 '2024-02-15', 0.5, 'ml', 'hydroponic', 'vegetative', 'controlled environment, 24°C', 'hydroponic solution, pH 5.8', 
 true, '2024-02-15', 0.5, 5, 'Hydroponic nutrients for optimal growth', '550e8400-e29b-41d4-a716-446655440000'),

('770e8400-e29b-41d4-a716-446655440010', '880e8400-e29b-41d4-a716-446655440010', '550e8400-e29b-41d4-a716-446655440010', 
 '2024-02-20', 0.3, 'g', 'fertigation', 'vegetative', 'controlled environment, 25°C', 'hydroponic solution, pH 6.0', 
 true, '2024-02-20', 0.3, 4, 'Calcium nitrate for strong plant structure', '550e8400-e29b-41d4-a716-446655440000'),

-- Future scheduled applications
('770e8400-e29b-41d4-a716-446655440011', '880e8400-e29b-41d4-a716-446655440011', '550e8400-e29b-41d4-a716-446655440003', 
 '2024-04-01', 0.1, 'g', 'foliar', 'fruiting', 'sunny, 30°C', 'loamy, pH 6.5', 
 false, NULL, NULL, NULL, 'Scheduled NPK application for fruit development', '550e8400-e29b-41d4-a716-446655440000'),

('770e8400-e29b-41d4-a716-446655440012', '880e8400-e29b-41d4-a716-446655440012', '550e8400-e29b-41d4-a716-446655440008', 
 '2024-04-05', 0.1, 'g', 'soil', 'pre_harvest', 'cloudy, 28°C', 'sandy loam, pH 6.2', 
 false, NULL, NULL, NULL, 'Final potassium application before harvest', '550e8400-e29b-41d4-a716-446655440000'),

-- Organic farming applications
('770e8400-e29b-41d4-a716-446655440013', '880e8400-e29b-41d4-a716-446655440013', '550e8400-e29b-41d4-a716-446655440001', 
 '2024-03-25', 1.0, 'kg', 'soil', 'vegetative', 'sunny, 26°C', 'organic soil, pH 6.7', 
 true, '2024-03-25', 1.0, 5, 'Organic compost for sustainable farming', '550e8400-e29b-41d4-a716-446655440000'),

('770e8400-e29b-41d4-a716-446655440014', '880e8400-e29b-41d4-a716-446655440014', '550e8400-e29b-41d4-a716-446655440005', 
 '2024-04-10', 0.02, 'ml', 'foliar', 'flowering', 'partly cloudy, 27°C', 'organic soil, pH 6.4', 
 false, NULL, NULL, NULL, 'Organic seaweed extract for flowering', '550e8400-e29b-41d4-a716-446655440000');
