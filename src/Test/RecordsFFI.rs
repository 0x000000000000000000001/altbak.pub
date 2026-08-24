
#[derive(Clone)]
struct Point {
    x: i64,
    y: i64,
    z: i64,
}

#[derive(Clone)]
struct DeepRecord {
    level1: Level1,
}

#[derive(Clone)]
struct Level1 {
    level2: Level2,
}

#[derive(Clone)]
struct Level2 {
    level3: Point,
}

pub fn Test_RecordsFFI_runRecordsFFI(mut limit: i64) -> i64 {
    let mut rec = DeepRecord {
        level1: Level1 {
            level2: Level2 {
                level3: Point { x: 0, y: 0, z: 0 },
            },
        },
    };
    for i in 0..limit {
        rec = DeepRecord {
            level1: Level1 {
                level2: Level2 {
                    level3: Point {
                        x: rec.level1.level2.level3.x + 1,
                        y: rec.level1.level2.level3.y + 2,
                        z: rec.level1.level2.level3.z + 3,
                    },
                },
            },
        };
    }
    rec.level1.level2.level3.x + rec.level1.level2.level3.y + rec.level1.level2.level3.z
}