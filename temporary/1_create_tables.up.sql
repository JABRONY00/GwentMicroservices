-- Players table
CREATE TABLE players (
    id VARCHAR(40) NOT NULL,
    name VARCHAR(20) UNIQUE NOT NULL,
    email VARCHAR(100) NOT NULL,
    password_hash BYTEA NOT NULL,
   -- profile_id SERIAL,

    PRIMARY KEY (id),
   -- FOREIGN KEY (profile_id) REFERENCES profiles(id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- players's game profile table
CREATE TABLE profiles (
    id VARCHAR(40) NOT NULL,
    wallet INT NOT NULL CHECK(wallet>=0) DEFAULT 300,
   active_race VARCHAR(4) NOT NULL,
   active_cards SMALLINT[],
   preset_id SERIAL,
    PRIMARY KEY (id),
    FOREIGN KEY (preset_id) REFERENCES presets(id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- players's presets
CREATE TABLE presets (
    id SERIAL,
    race VARCHAR(10) NOT NULL,

    department_id INT,
    PRIMARY KEY (id),
    FOREIGN KEY (department_id) REFERENCES departments(id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- Связующая таблица между сотрудниками и кабинетами (многие ко многим)
CREATE TABLE employee_rooms (
    id SERIAL,
    employee_id VARCHAR(40),
    room_id INT,
    PRIMARY KEY (id),
    FOREIGN KEY (employee_id) REFERENCES employees(id),
    FOREIGN KEY (room_id) REFERENCES rooms(id)
);

-- Создание таблицы расходных материалов в кабинетах
CREATE TABLE room_storage (
    id SERIAL,
    item_name VARCHAR(100) NOT NULL,
    quantity INT NOT NULL,
    room_id INT,
    PRIMARY KEY (id),
    FOREIGN KEY (room_id) REFERENCES rooms(id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- Создание таблицы заявок на посещение врача
CREATE TABLE appointments (
    id SERIAL,
    appointment_date DATE NOT NULL,
    appointment_time TIME NOT NULL,
    status VARCHAR(50),

    client_id VARCHAR(40),
    employee_id VARCHAR(40),
    PRIMARY KEY (id),
    FOREIGN KEY (client_id) REFERENCES clients(id),
    FOREIGN KEY (employee_id) REFERENCES employees(id)
);

-- Создание таблицы посещений врача
CREATE TABLE medical_sessions (
    id SERIAL,
    session_date DATE NOT NULL,
    session_time TIME NOT NULL,
    appointment_date DATE NOT NULL,
    appointment_time TIME NOT NULL,
    status VARCHAR(50),
    comments TEXT,
    attached_files TEXT, -- Пути к файлам

    client_id VARCHAR(40),
    employee_id VARCHAR(40),
    PRIMARY KEY (id),
    FOREIGN KEY (client_id) REFERENCES clients(id),
    FOREIGN KEY (employee_id) REFERENCES employees(id)
);

-- Создание таблицы медицинских карт
CREATE TABLE medical_cards (
    id SERIAL,
    health_info TEXT NOT NULL,
    last_updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    client_id VARCHAR(40),
    PRIMARY KEY (id),
    FOREIGN KEY (client_id) REFERENCES clients(id)
);

-- Создание таблицы смен сотрудников
CREATE TABLE shifts (
    id SERIAL,
    shift_date DATE NOT NULL,
    start_time TIME NOT NULL,
    end_time TIME NOT NULL,

    employee_id VARCHAR(40),
    PRIMARY KEY (id),
    FOREIGN KEY (employee_id) REFERENCES employees(id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- Создание таблицы выходных, отпусков и больничных сотрудников
CREATE TABLE time_off (
    id SERIAL,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    type VARCHAR(50) NOT NULL, -- Тип (отпуск, больничный и т.д.)

    employee_id VARCHAR(40),
    PRIMARY KEY (id),
    FOREIGN KEY (employee_id) REFERENCES employees(id) ON DELETE CASCADE ON UPDATE CASCADE
);

