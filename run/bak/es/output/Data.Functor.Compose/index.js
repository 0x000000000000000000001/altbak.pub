const Compose = x => x;
const showCompose = dictShow => ({show: v => "(Compose " + dictShow.show(v) + ")"});
const newtypeCompose = {Coercible0: () => {}};
const functorCompose = dictFunctor => dictFunctor1 => ({map: f => v => dictFunctor.map(dictFunctor1.map(f))(v)});
const eqCompose = dictEq1 => dictEq11 => dictEq => {
  const eqApp1 = {eq: x => y => dictEq11.eq1(dictEq)(x)(y)};
  return {eq: v => v1 => dictEq1.eq1(eqApp1)(v)(v1)};
};
const ordCompose = dictOrd1 => {
  const $0 = dictOrd1.Eq10();
  return dictOrd11 => {
    const $1 = dictOrd11.Eq10();
    const $2 = dictOrd11.Eq10();
    return dictOrd => {
      const $3 = dictOrd.Eq0();
      const eqApp2 = {eq: x => y => $1.eq1($3)(x)(y)};
      const ordApp1 = {compare: x => y => dictOrd11.compare1(dictOrd)(x)(y), Eq0: () => eqApp2};
      const $4 = dictOrd.Eq0();
      const eqApp1 = {eq: x => y => $2.eq1($4)(x)(y)};
      const eqCompose3 = {eq: v => v1 => $0.eq1(eqApp1)(v)(v1)};
      return {compare: v => v1 => dictOrd1.compare1(ordApp1)(v)(v1), Eq0: () => eqCompose3};
    };
  };
};
const eq1Compose = dictEq1 => dictEq11 => (
  {
    eq1: dictEq => {
      const eqApp1 = {eq: x => y => dictEq11.eq1(dictEq)(x)(y)};
      return v => v1 => dictEq1.eq1(eqApp1)(v)(v1);
    }
  }
);
const ord1Compose = dictOrd1 => {
  const ordCompose1 = ordCompose(dictOrd1);
  const $0 = dictOrd1.Eq10();
  return dictOrd11 => {
    const ordCompose2 = ordCompose1(dictOrd11);
    const $1 = dictOrd11.Eq10();
    const eq1Compose2 = {
      eq1: dictEq => {
        const eqApp1 = {eq: x => y => $1.eq1(dictEq)(x)(y)};
        return v => v1 => $0.eq1(eqApp1)(v)(v1);
      }
    };
    return {compare1: dictOrd => ordCompose2(dictOrd).compare, Eq10: () => eq1Compose2};
  };
};
const bihoistCompose = dictFunctor => natF => natG => v => natF(dictFunctor.map(natG)(v));
const applyCompose = dictApply => {
  const Functor0 = dictApply.Functor0();
  const $0 = dictApply.Functor0();
  return dictApply1 => {
    const apply = dictApply1.apply;
    const $1 = dictApply1.Functor0();
    const functorCompose2 = {map: f => v => $0.map($1.map(f))(v)};
    return {apply: v => v1 => dictApply.apply(Functor0.map(apply)(v))(v1), Functor0: () => functorCompose2};
  };
};
const applicativeCompose = dictApplicative => {
  const $0 = dictApplicative.Apply0();
  const Functor0 = $0.Functor0();
  const applyCompose1 = (() => {
    const $1 = $0.Functor0();
    return dictApply1 => {
      const apply = dictApply1.apply;
      const $2 = dictApply1.Functor0();
      const functorCompose2 = {map: f => v => $1.map($2.map(f))(v)};
      return {apply: v => v1 => $0.apply(Functor0.map(apply)(v))(v1), Functor0: () => functorCompose2};
    };
  })();
  return dictApplicative1 => {
    const applyCompose2 = applyCompose1(dictApplicative1.Apply0());
    return {pure: x => dictApplicative.pure(dictApplicative1.pure(x)), Apply0: () => applyCompose2};
  };
};
const altCompose = dictAlt => {
  const $0 = dictAlt.Functor0();
  return dictFunctor => {
    const functorCompose2 = {map: f => v => $0.map(dictFunctor.map(f))(v)};
    return {alt: v => v1 => dictAlt.alt(v)(v1), Functor0: () => functorCompose2};
  };
};
const plusCompose = dictPlus => {
  const empty = dictPlus.empty;
  const $0 = dictPlus.Alt0();
  const $1 = $0.Functor0();
  return dictFunctor => {
    const functorCompose2 = {map: f => v => $1.map(dictFunctor.map(f))(v)};
    const altCompose2 = {alt: v => v1 => $0.alt(v)(v1), Functor0: () => functorCompose2};
    return {empty, Alt0: () => altCompose2};
  };
};
const alternativeCompose = dictAlternative => {
  const applicativeCompose1 = applicativeCompose(dictAlternative.Applicative0());
  const $0 = dictAlternative.Plus1();
  const empty = $0.empty;
  const plusCompose1 = (() => {
    const $1 = $0.Alt0();
    const $2 = $1.Functor0();
    return dictFunctor => {
      const functorCompose2 = {map: f => v => $2.map(dictFunctor.map(f))(v)};
      const altCompose2 = {alt: v => v1 => $1.alt(v)(v1), Functor0: () => functorCompose2};
      return {empty, Alt0: () => altCompose2};
    };
  })();
  return dictApplicative => {
    const applicativeCompose2 = applicativeCompose1(dictApplicative);
    const plusCompose2 = plusCompose1(dictApplicative.Apply0().Functor0());
    return {Applicative0: () => applicativeCompose2, Plus1: () => plusCompose2};
  };
};
export {
  Compose,
  altCompose,
  alternativeCompose,
  applicativeCompose,
  applyCompose,
  bihoistCompose,
  eq1Compose,
  eqCompose,
  functorCompose,
  newtypeCompose,
  ord1Compose,
  ordCompose,
  plusCompose,
  showCompose
};
