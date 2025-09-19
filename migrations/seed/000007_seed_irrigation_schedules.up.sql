-- Seed data for irrigation_schedules table
INSERT INTO irrigation_schedules (
    id, growing_zone_id, planting_cycle_id, schedule_name, irrigation_type, 
    start_time, duration_minutes, frequency, days_of_week, water_amount_liters, 
    fertilizer_mix, is_active, created_by
) VALUES 
-- Morning irrigation schedules
('660e8400-e29b-41d4-a716-446655440001', '770e8400-e29b-41d4-a716-446655440001', '880e8400-e29b-41d4-a716-446655440001', 
 'Morning Drip Irrigation - Zone A', 'drip', '06:00:00', 30, 'daily', 
 '["monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"]', 
 50.0, false, true, '550e8400-e29b-41d4-a716-446655440000'),

('660e8400-e29b-41d4-a716-446655440002', '770e8400-e29b-41d4-a716-446655440002', '880e8400-e29b-41d4-a716-446655440002', 
 'Morning Spray Irrigation - Zone B', 'spray', '07:00:00', 45, 'daily', 
 '["monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"]', 
 75.0, false, true, '550e8400-e29b-41d4-a716-446655440000'),

-- Evening irrigation schedules
('660e8400-e29b-41d4-a716-446655440003', '770e8400-e29b-41d4-a716-446655440003', '880e8400-e29b-41d4-a716-446655440003', 
 'Evening Drip Irrigation - Zone C', 'drip', '18:00:00', 25, 'daily', 
 '["monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"]', 
 40.0, false, true, '550e8400-e29b-41d4-a716-446655440000'),

-- Weekly irrigation schedules
('660e8400-e29b-41d4-a716-446655440004', '770e8400-e29b-41d4-a716-446655440004', '880e8400-e29b-41d4-a716-446655440004', 
 'Weekly Flood Irrigation - Zone D', 'flood', '08:00:00', 60, 'weekly', 
 '["monday"]', 200.0, false, true, '550e8400-e29b-41d4-a716-446655440000'),

-- Fertigation schedules (irrigation with fertilizer)
('660e8400-e29b-41d4-a716-446655440005', '770e8400-e29b-41d4-a716-446655440005', '880e8400-e29b-41d4-a716-446655440005', 
 'Morning Fertigation - Zone E', 'drip', '06:30:00', 40, 'bi_weekly', 
 '["monday", "thursday"]', 60.0, true, true, '550e8400-e29b-41d4-a716-446655440000'),

('660e8400-e29b-41d4-a716-446655440006', '770e8400-e29b-41d4-a716-446655440006', '880e8400-e29b-41d4-a716-446655440006', 
 'Evening Fertigation - Zone F', 'drip', '17:30:00', 35, 'bi_weekly', 
 '["tuesday", "friday"]', 55.0, true, true, '550e8400-e29b-41d4-a716-446655440000'),

-- Custom irrigation schedules
('660e8400-e29b-41d4-a716-446655440007', '770e8400-e29b-41d4-a716-446655440007', '880e8400-e29b-41d4-a716-446655440007', 
 'Custom Schedule - Zone G', 'automatic', '09:00:00', 20, 'custom', 
 '["monday", "wednesday", "friday"]', 30.0, false, true, '550e8400-e29b-41d4-a716-446655440000'),

-- Manual irrigation schedules
('660e8400-e29b-41d4-a716-446655440008', '770e8400-e29b-41d4-a716-446655440008', '880e8400-e29b-41d4-a716-446655440008', 
 'Manual Irrigation - Zone H', 'manual', '10:00:00', 0, 'manual', 
 '[]', 0.0, false, false, '550e8400-e29b-41d4-a716-446655440000'),

-- High-frequency irrigation for seedlings
('660e8400-e29b-41d4-a716-446655440009', '770e8400-e29b-41d4-a716-446655440009', '880e8400-e29b-41d4-a716-446655440009', 
 'Seedling Care - Zone I', 'drip', '08:00:00', 15, 'daily', 
 '["monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"]', 
 20.0, false, true, '550e8400-e29b-41d4-a716-446655440000'),

-- Drought-resistant crop irrigation
('660e8400-e29b-41d4-a716-446655440010', '770e8400-e29b-41d4-a716-446655440010', '880e8400-e29b-41d4-a716-446655440010', 
 'Drought Schedule - Zone J', 'drip', '05:00:00', 50, 'weekly', 
 '["sunday"]', 100.0, false, true, '550e8400-e29b-41d4-a716-446655440000');
