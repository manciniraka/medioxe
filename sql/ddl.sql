CREATE TABLE IF NOT EXISTS users (
    id  SERIAL PRIMARY KEY,
    full_name   VARCHAR(100) NOT NULL,
    email   VARCHAR(100) NOT NULL UNIQUE,
    password    VARCHAR(255) NOT NULL,
    role    VARCHAR(50) NOT NULL,
    created_at  TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS specialties (
    id  SERIAL PRIMARY KEY,
    name   VARCHAR(100) NOT NULL,
    created_at  TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS hospitals (
    id  SERIAL PRIMARY KEY,
    name   VARCHAR(100) NOT NULL,
    address TEXT NOT NULL.
    created_at  TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS doctor_profiles (
    id  SERIAL PRIMARY KEY,
    user_id INT UNIQUE NOT NULL REFERENCES users(id),
    specialty_id INT NOT NULL REFERENCES specialties(id),
    hospital_id INT NOT NULL REFERENCES hospitals(id),
    experience_years   INT NOT NULL,
    consultation_fee   INT NOT NULL,
    bio    TEXT NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at  TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS schedules (
    id  SERIAL PRIMARY KEY,
    doctor_id INT NOT NULL REFERENCES doctor_profiles(id),
    date  DATE NOT NULL,
    start_time  TIME NOT NULL,
    end_time  TIME NOT NULL,
    is_booked    BOOLEAN NOT NULL DEFAULT FALSE,
    created_at  TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS symptom_analyses (
    id  SERIAL PRIMARY KEY,
    patient_id INT NOT NULL REFERENCES users(id),
    symptoms TEXT NOT NULL,
    recommended_specialty_id INT NOT NULL REFERENCES specialties(id),
    recommended_doctor_id INT REFERENCES doctor_profiles(id),
    ai_summary    TEXT NOT NULL,
    created_at  TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS appointments (
    id  SERIAL PRIMARY KEY,
    patient_id INT NOT NULL REFERENCES users(id),
    doctor_id INT NOT NULL REFERENCES doctor_profiles(id),
    schedule_id INT NOT NULL REFERENCES schedules(id),
    symptom_analysis_id INT REFERENCES symptom_analyses(id),
    notes TEXT,
    status VARCHAR(50) NOT NULL,
    created_at  TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS appointment_histories (
    id  SERIAL PRIMARY KEY,
    appointment_id INT NOT NULL REFERENCES appointments(id),
    status VARCHAR(50) NOT NULL,
    remark TEXT,
    created_at  TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS payments (
    id  SERIAL PRIMARY KEY,
    appointment_id INT NOT NULL REFERENCES appointments(id),
    amount INT NOT NULL,
    status VARCHAR(50) NOT NULL,
    transaction_id VARCHAR(255),
    paid_at  TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at  TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP NOT NULL DEFAULT NOW()
);