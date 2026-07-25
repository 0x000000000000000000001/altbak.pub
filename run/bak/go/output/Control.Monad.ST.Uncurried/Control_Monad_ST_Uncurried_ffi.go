package Control_Monad_ST_Uncurried

import "gopurs/output/gopurs_runtime"

func MkSTFn1(fn any) any {
	return func(a any) any {
		return fn.(func(any) any)(a).(func(any) any)(nil)
	}
}

func MkSTFn2(fn any) any {
	return func(a any) any {
		return func(b any) any {
			return fn.(func(any) any)(a).(func(any) any)(b).(func(any) any)(nil)
		}
	}
}

func MkSTFn3(fn any) any {
	return func(a any) any {
		return func(b any) any {
			return func(c any) any {
				return fn.(func(any) any)(a).(func(any) any)(b).(func(any) any)(c).(func(any) any)(nil)
			}
		}
	}
}

func MkSTFn4(fn any) any {
	return func(a any) any {
		return func(b any) any {
			return func(c any) any {
				return func(d any) any {
					return fn.(func(any) any)(a).(func(any) any)(b).(func(any) any)(c).(func(any) any)(d).(func(any) any)(nil)
				}
			}
		}
	}
}

func MkSTFn5(fn any) any {
	return func(a any) any {
		return func(b any) any {
			return func(c any) any {
				return func(d any) any {
					return func(e any) any {
						return fn.(func(any) any)(a).(func(any) any)(b).(func(any) any)(c).(func(any) any)(d).(func(any) any)(e).(func(any) any)(nil)
					}
				}
			}
		}
	}
}

func MkSTFn6(fn any) any {
	return func(a any) any {
		return func(b any) any {
			return func(c any) any {
				return func(d any) any {
					return func(e any) any {
						return func(f any) any {
							return fn.(func(any) any)(a).(func(any) any)(b).(func(any) any)(c).(func(any) any)(d).(func(any) any)(e).(func(any) any)(f).(func(any) any)(nil)
						}
					}
				}
			}
		}
	}
}

func MkSTFn7(fn any) any {
	return func(a any) any {
		return func(b any) any {
			return func(c any) any {
				return func(d any) any {
					return func(e any) any {
						return func(f any) any {
							return func(g any) any {
								return fn.(func(any) any)(a).(func(any) any)(b).(func(any) any)(c).(func(any) any)(d).(func(any) any)(e).(func(any) any)(f).(func(any) any)(g).(func(any) any)(nil)
							}
						}
					}
				}
			}
		}
	}
}

func MkSTFn8(fn any) any {
	return func(a any) any {
		return func(b any) any {
			return func(c any) any {
				return func(d any) any {
					return func(e any) any {
						return func(f any) any {
							return func(g any) any {
								return func(h any) any {
									return fn.(func(any) any)(a).(func(any) any)(b).(func(any) any)(c).(func(any) any)(d).(func(any) any)(e).(func(any) any)(f).(func(any) any)(g).(func(any) any)(h).(func(any) any)(nil)
								}
							}
						}
					}
				}
			}
		}
	}
}

func MkSTFn9(fn any) any {
	return func(a any) any {
		return func(b any) any {
			return func(c any) any {
				return func(d any) any {
					return func(e any) any {
						return func(f any) any {
							return func(g any) any {
								return func(h any) any {
									return func(i any) any {
										return fn.(func(any) any)(a).(func(any) any)(b).(func(any) any)(c).(func(any) any)(d).(func(any) any)(e).(func(any) any)(f).(func(any) any)(g).(func(any) any)(h).(func(any) any)(i).(func(any) any)(nil)
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

func MkSTFn10(fn any) any {
	return func(a any) any {
		return func(b any) any {
			return func(c any) any {
				return func(d any) any {
					return func(e any) any {
						return func(f any) any {
							return func(g any) any {
								return func(h any) any {
									return func(i any) any {
										return func(j any) any {
											return fn.(func(any) any)(a).(func(any) any)(b).(func(any) any)(c).(func(any) any)(d).(func(any) any)(e).(func(any) any)(f).(func(any) any)(g).(func(any) any)(h).(func(any) any)(i).(func(any) any)(j).(func(any) any)(nil)
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

func RunSTFn1(fn any) any {
	return func(a any) any {
		return func(_ any) any {
			return fn.(func(any) any)(a)
		}
	}
}

func RunSTFn2(fn any) any {
	return func(a any) any {
		return func(b any) any {
			return func(_ any) any {
				return fn.(func(any) any)(a).(func(any) any)(b)
			}
		}
	}
}

func RunSTFn3(fn any) any {
	return func(a any) any {
		return func(b any) any {
			return func(c any) any {
				return func(_ any) any {
					return fn.(func(any) any)(a).(func(any) any)(b).(func(any) any)(c)
				}
			}
		}
	}
}

func RunSTFn4(fn any) any {
	return func(a any) any {
		return func(b any) any {
			return func(c any) any {
				return func(d any) any {
					return func(_ any) any {
						return fn.(func(any) any)(a).(func(any) any)(b).(func(any) any)(c).(func(any) any)(d)
					}
				}
			}
		}
	}
}

func RunSTFn5(fn any) any {
	return func(a any) any {
		return func(b any) any {
			return func(c any) any {
				return func(d any) any {
					return func(e any) any {
						return func(_ any) any {
							return fn.(func(any) any)(a).(func(any) any)(b).(func(any) any)(c).(func(any) any)(d).(func(any) any)(e)
						}
					}
				}
			}
		}
	}
}

func RunSTFn6(fn any) any {
	return func(a any) any {
		return func(b any) any {
			return func(c any) any {
				return func(d any) any {
					return func(e any) any {
						return func(f any) any {
							return func(_ any) any {
								return fn.(func(any) any)(a).(func(any) any)(b).(func(any) any)(c).(func(any) any)(d).(func(any) any)(e).(func(any) any)(f)
							}
						}
					}
				}
			}
		}
	}
}

func RunSTFn7(fn any) any {
	return func(a any) any {
		return func(b any) any {
			return func(c any) any {
				return func(d any) any {
					return func(e any) any {
						return func(f any) any {
							return func(g any) any {
								return func(_ any) any {
									return fn.(func(any) any)(a).(func(any) any)(b).(func(any) any)(c).(func(any) any)(d).(func(any) any)(e).(func(any) any)(f).(func(any) any)(g)
								}
							}
						}
					}
				}
			}
		}
	}
}

func RunSTFn8(fn any) any {
	return func(a any) any {
		return func(b any) any {
			return func(c any) any {
				return func(d any) any {
					return func(e any) any {
						return func(f any) any {
							return func(g any) any {
								return func(h any) any {
									return func(_ any) any {
										return fn.(func(any) any)(a).(func(any) any)(b).(func(any) any)(c).(func(any) any)(d).(func(any) any)(e).(func(any) any)(f).(func(any) any)(g).(func(any) any)(h)
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

func RunSTFn9(fn any) any {
	return func(a any) any {
		return func(b any) any {
			return func(c any) any {
				return func(d any) any {
					return func(e any) any {
						return func(f any) any {
							return func(g any) any {
								return func(h any) any {
									return func(i any) any {
										return func(_ any) any {
											return fn.(func(any) any)(a).(func(any) any)(b).(func(any) any)(c).(func(any) any)(d).(func(any) any)(e).(func(any) any)(f).(func(any) any)(g).(func(any) any)(h).(func(any) any)(i)
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

func RunSTFn10(fn any) any {
	return func(a any) any {
		return func(b any) any {
			return func(c any) any {
				return func(d any) any {
					return func(e any) any {
						return func(f any) any {
							return func(g any) any {
								return func(h any) any {
									return func(i any) any {
										return func(j any) any {
											return func(_ any) any {
												return fn.(func(any) any)(a).(func(any) any)(b).(func(any) any)(c).(func(any) any)(d).(func(any) any)(e).(func(any) any)(f).(func(any) any)(g).(func(any) any)(h).(func(any) any)(i).(func(any) any)(j)
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}


// --- Auto-generated FFI wrappers ---
func Call_mkSTFn1(arg0 any) any {
	return MkSTFn1(arg0)
}
var _Gopurs_MkSTFn1 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MkSTFn1(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_mkSTFn2(arg0 any) any {
	return MkSTFn2(arg0)
}
var _Gopurs_MkSTFn2 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MkSTFn2(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_mkSTFn3(arg0 any) any {
	return MkSTFn3(arg0)
}
var _Gopurs_MkSTFn3 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MkSTFn3(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_mkSTFn4(arg0 any) any {
	return MkSTFn4(arg0)
}
var _Gopurs_MkSTFn4 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MkSTFn4(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_mkSTFn5(arg0 any) any {
	return MkSTFn5(arg0)
}
var _Gopurs_MkSTFn5 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MkSTFn5(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_mkSTFn6(arg0 any) any {
	return MkSTFn6(arg0)
}
var _Gopurs_MkSTFn6 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MkSTFn6(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_mkSTFn7(arg0 any) any {
	return MkSTFn7(arg0)
}
var _Gopurs_MkSTFn7 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MkSTFn7(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_mkSTFn8(arg0 any) any {
	return MkSTFn8(arg0)
}
var _Gopurs_MkSTFn8 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MkSTFn8(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_mkSTFn9(arg0 any) any {
	return MkSTFn9(arg0)
}
var _Gopurs_MkSTFn9 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MkSTFn9(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_mkSTFn10(arg0 any) any {
	return MkSTFn10(arg0)
}
var _Gopurs_MkSTFn10 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MkSTFn10(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_runSTFn1(arg0 any) any {
	return RunSTFn1(arg0)
}
var _Gopurs_RunSTFn1 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := RunSTFn1(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_runSTFn2(arg0 any) any {
	return RunSTFn2(arg0)
}
var _Gopurs_RunSTFn2 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := RunSTFn2(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_runSTFn3(arg0 any) any {
	return RunSTFn3(arg0)
}
var _Gopurs_RunSTFn3 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := RunSTFn3(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_runSTFn4(arg0 any) any {
	return RunSTFn4(arg0)
}
var _Gopurs_RunSTFn4 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := RunSTFn4(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_runSTFn5(arg0 any) any {
	return RunSTFn5(arg0)
}
var _Gopurs_RunSTFn5 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := RunSTFn5(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_runSTFn6(arg0 any) any {
	return RunSTFn6(arg0)
}
var _Gopurs_RunSTFn6 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := RunSTFn6(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_runSTFn7(arg0 any) any {
	return RunSTFn7(arg0)
}
var _Gopurs_RunSTFn7 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := RunSTFn7(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_runSTFn8(arg0 any) any {
	return RunSTFn8(arg0)
}
var _Gopurs_RunSTFn8 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := RunSTFn8(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_runSTFn9(arg0 any) any {
	return RunSTFn9(arg0)
}
var _Gopurs_RunSTFn9 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := RunSTFn9(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_runSTFn10(arg0 any) any {
	return RunSTFn10(arg0)
}
var _Gopurs_RunSTFn10 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := RunSTFn10(go_arg0)
	return gopurs_runtime.Box(go_res)
})
