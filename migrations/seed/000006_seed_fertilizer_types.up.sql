-- Seed data for fertilizer_types table
INSERT INTO fertilizer_types (
    id, name, type, npk_ratio, nitrogen_percentage, phosphorus_percentage, 
    potassium_percentage, trace_elements, application_method, dosage_per_plant, 
    dosage_per_m2, unit, manufacturer, cost_per_unit, description, 
    safety_notes, status, created_by
) VALUES 
-- Organic fertilizers
('550e8400-e29b-41d4-a716-446655440001', 'Compost', 'organic', '2-1-1', 2.0, 1.0, 1.0, 
 '{"calcium": 0.5, "magnesium": 0.3, "iron": 0.1}', 'soil', 0.5, 2.0, 'kg', 
 'Local Farm', 15000.00, 'Natural compost from organic waste', 
 'Safe for organic farming', 'active', '550e8400-e29b-41d4-a716-446655440000'),

('550e8400-e29b-41d4-a716-446655440002', 'Chicken Manure', 'organic', '3-2-1', 3.0, 2.0, 1.0, 
 '{"calcium": 1.2, "magnesium": 0.8}', 'soil', 0.3, 1.5, 'kg', 
 'Poultry Farm Co.', 12000.00, 'Dried and processed chicken manure', 
 'Must be composted before use', 'active', '550e8400-e29b-41d4-a716-446655440000'),

-- Chemical fertilizers
('550e8400-e29b-41d4-a716-446655440003', 'NPK 20-20-20', 'chemical', '20-20-20', 20.0, 20.0, 20.0, 
 '{"sulfur": 2.0, "calcium": 1.0}', 'foliar', 0.1, 0.5, 'g', 
 'AgroChem Ltd.', 25000.00, 'Balanced NPK fertilizer for general use', 
 'Use protective equipment when applying', 'active', '550e8400-e29b-41d4-a716-446655440000'),

('550e8400-e29b-41d4-a716-446655440004', 'Urea', 'chemical', '46-0-0', 46.0, 0.0, 0.0, 
 '{"sulfur": 0.5}', 'soil', 0.05, 0.2, 'g', 
 'Nitrogen Corp.', 18000.00, 'High nitrogen fertilizer for vegetative growth', 
 'Avoid contact with skin and eyes', 'active', '550e8400-e29b-41d4-a716-446655440000'),

-- Liquid fertilizers
('550e8400-e29b-41d4-a716-446655440005', 'Liquid Seaweed', 'liquid', '1-1-1', 1.0, 1.0, 1.0, 
 '{"iodine": 0.1, "zinc": 0.05, "manganese": 0.02}', 'foliar', 0.02, 0.1, 'ml', 
 'OceanGrow Inc.', 35000.00, 'Natural seaweed extract with trace elements', 
 'Dilute before application', 'active', '550e8400-e29b-41d4-a716-446655440000'),

('550e8400-e29b-41d4-a716-446655440006', 'Fish Emulsion', 'liquid', '5-1-1', 5.0, 1.0, 1.0, 
 '{"calcium": 0.8, "magnesium": 0.3}', 'soil', 0.1, 0.5, 'ml', 
 'AquaGrow Ltd.', 28000.00, 'Liquid fish fertilizer for organic farming', 
 'Store in cool place, may have strong odor', 'active', '550e8400-e29b-41d4-a716-446655440000'),

-- Granular fertilizers
('550e8400-e29b-41d4-a716-446655440007', 'Superphosphate', 'granular', '0-20-0', 0.0, 20.0, 0.0, 
 '{"calcium": 15.0, "sulfur": 3.0}', 'soil', 0.2, 1.0, 'g', 
 'Phosphate Corp.', 22000.00, 'High phosphorus fertilizer for root development', 
 'Keep away from children and pets', 'active', '550e8400-e29b-41d4-a716-446655440000'),

('550e8400-e29b-41d4-a716-446655440008', 'Potassium Sulfate', 'granular', '0-0-50', 0.0, 0.0, 50.0, 
 '{"sulfur": 18.0}', 'soil', 0.1, 0.5, 'g', 
 'Potash Ltd.', 30000.00, 'High potassium fertilizer for flowering and fruiting', 
 'Apply evenly to avoid burning', 'active', '550e8400-e29b-41d4-a716-446655440000'),

-- Hydroponic fertilizers
('550e8400-e29b-41d4-a716-446655440009', 'Hydroponic A+B', 'liquid', '5-3-4', 5.0, 3.0, 4.0, 
 '{"calcium": 2.0, "magnesium": 1.0, "iron": 0.1, "manganese": 0.05}', 'hydroponic', 0.5, 2.0, 'ml', 
 'HydroGrow Systems', 45000.00, 'Complete nutrient solution for hydroponic systems', 
 'Mix parts A and B separately before combining', 'active', '550e8400-e29b-41d4-a716-446655440000'),

('550e8400-e29b-41d4-a716-446655440010', 'Calcium Nitrate', 'powder', '15-0-0', 15.0, 0.0, 0.0, 
 '{"calcium": 19.0}', 'fertigation', 0.3, 1.5, 'g', 
 'CalNit Industries', 20000.00, 'Calcium and nitrogen source for fertigation', 
 'Store in dry place, avoid moisture', 'active', '550e8400-e29b-41d4-a716-446655440000');
