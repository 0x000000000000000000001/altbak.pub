#![allow(warnings)]
#[path="tracked_rc.rs"] mod tracked;
include!("kernel-count-borrow-ins-hoisted.rs");
include!("count_harness.rs");
