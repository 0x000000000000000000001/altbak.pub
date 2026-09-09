fn main(){
 for seed in [0_i64,1,7,50] {
  for n in [0_i64,1,2,5,10,10000] {
   let r=Test_Records_updateRec(seed,Test_Records_initial());
   let old=values(&r);
   let kept=r.clone();
   let old_b=r.get_b();let old_d=old_b.get_d();
   let result=Test_Records_updateRec(n,r);
   assert_eq!(values(&result),expected(n,old));assert_eq!(values(&kept),old);
   assert_eq!(old_b.get_c().unwrap_int(),old[1]);assert_eq!(old_d.get_e().unwrap_int(),old[2]);
   assert_eq!(old_d.get_f().unwrap_int(),old[3]);
   // Sharing a nested node alone must also retain its original version.
   let r=Test_Records_updateRec(seed,Test_Records_initial());
   let kept_leaf=r.get_b().get_d();
   let result=Test_Records_updateRec(n,r);
   assert_eq!(values(&result),expected(n,old));
   assert_eq!(kept_leaf.get_e().unwrap_int(),old[2]);
   assert_eq!(kept_leaf.get_f().unwrap_int(),old[3]);
  }
 }
 for mask in 0..8 {
 for seed in [0_i64,7] {
  let r=Test_Records_updateRec(seed,Test_Records_initial());
  let old=values(&r);
  let root=if mask&1!=0 {Some(r.clone())} else {None};
  let middle=if mask&2!=0 {Some(r.get_b())} else {None};
  let leaf=if mask&4!=0 {Some(r.get_b().get_d())} else {None};
  let result=Test_Records_updateRec(10000,r);
  assert_eq!(values(&result),expected(10000,old));
  if let Some(r)=root {assert_eq!(values(&r),old);}
  if let Some(b)=middle {assert_eq!(b.get_c().unwrap_int(),old[1]); assert_eq!(b.get_d().get_e().unwrap_int(),old[2]); assert_eq!(b.get_d().get_f().unwrap_int(),old[3]);}
  if let Some(d)=leaf {assert_eq!(d.get_e().unwrap_int(),old[2]); assert_eq!(d.get_f().unwrap_int(),old[3]);}
 }
}
 println!("8 independent sharing masks and 24 seed/count combinations: all four fields, retained roots, nested and leaf-only sharing passed");
}
