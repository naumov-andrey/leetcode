# Write your MySQL query statement below

SELECT Weather.id
FROM Weather
JOIN Weather AS yesterday ON yesterday.recordDate = SUBDATE(Weather.recordDate, 1)
WHERE Weather.temperature > yesterday.temperature 
