package Effect_Uncurried

import "gopurs/output/gopurs_runtime"


func MkEffectFn1(f any) func(any) func() any {
	return func(a any) func() any {
		return func() any {
			return f.(func(any) any)(a)
		}
	}
}

func MkEffectFn2(f any) func(any, any) func() any {
	return func(a, b any) func() any {
		return func() any {
			return f.(func(any) any)(a).(func(any) any)(b)
		}
	}
}

func MkEffectFn3(f any) func(any, any, any) func() any {
	return func(a, b, c any) func() any {
		return func() any {
			return f.(func(any) any)(a).(func(any) any)(b).(func(any) any)(c)
		}
	}
}

func MkEffectFn4(f any) func(any, any, any, any) func() any {
	return func(a, b, c, d any) func() any {
		return func() any {
			return f.(func(any) any)(a).(func(any) any)(b).(func(any) any)(c).(func(any) any)(d)
		}
	}
}

func MkEffectFn5(f any) func(any, any, any, any, any) func() any {
	return func(a, b, c, d, e any) func() any {
		return func() any {
			return f.(func(any) any)(a).(func(any) any)(b).(func(any) any)(c).(func(any) any)(d).(func(any) any)(e)
		}
	}
}

func MkEffectFn6(f any) func(any, any, any, any, any, any) func() any {
	return func(a, b, c, d, e, g any) func() any {
		return func() any {
			return f.(func(any) any)(a).(func(any) any)(b).(func(any) any)(c).(func(any) any)(d).(func(any) any)(e).(func(any) any)(g)
		}
	}
}

func MkEffectFn7(f any) func(any, any, any, any, any, any, any) func() any {
	return func(a, b, c, d, e, g, h any) func() any {
		return func() any {
			return f.(func(any) any)(a).(func(any) any)(b).(func(any) any)(c).(func(any) any)(d).(func(any) any)(e).(func(any) any)(g).(func(any) any)(h)
		}
	}
}

func MkEffectFn8(f any) func(any, any, any, any, any, any, any, any) func() any {
	return func(a, b, c, d, e, g, h, i any) func() any {
		return func() any {
			return f.(func(any) any)(a).(func(any) any)(b).(func(any) any)(c).(func(any) any)(d).(func(any) any)(e).(func(any) any)(g).(func(any) any)(h).(func(any) any)(i)
		}
	}
}

func MkEffectFn9(f any) func(any, any, any, any, any, any, any, any, any) func() any {
	return func(a, b, c, d, e, g, h, i, j any) func() any {
		return func() any {
			return f.(func(any) any)(a).(func(any) any)(b).(func(any) any)(c).(func(any) any)(d).(func(any) any)(e).(func(any) any)(g).(func(any) any)(h).(func(any) any)(i).(func(any) any)(j)
		}
	}
}

func MkEffectFn10(f any) func(any, any, any, any, any, any, any, any, any, any) func() any {
	return func(a, b, c, d, e, g, h, i, j, k any) func() any {
		return func() any {
			return f.(func(any) any)(a).(func(any) any)(b).(func(any) any)(c).(func(any) any)(d).(func(any) any)(e).(func(any) any)(g).(func(any) any)(h).(func(any) any)(i).(func(any) any)(j).(func(any) any)(k)
		}
	}
}

func RunEffectFn1(f any) func(any) func() any {
	return func(a any) func() any {
		return func() any {
			return f.(func(any) any)(a)
		}
	}
}

func RunEffectFn2(f any) func(any) func(any) func() any {
	return func(a any) func(any) func() any {
		return func(b any) func() any {
			return func() any {
				return f.(func(any) any)(a, b)
			}
		}
	}
}

func RunEffectFn3(f any) func(any) func(any) func(any) func() any {
	return func(a any) func(any) func(any) func() any {
		return func(b any) func(any) func() any {
			return func(c any) func() any {
				return func() any {
					return f.(func(any) any)(a, b, c)
				}
			}
		}
	}
}

func RunEffectFn4(f any) func(any) func(any) func(any) func(any) func() any {
	return func(a any) func(any) func(any) func(any) func() any {
		return func(b any) func(any) func(any) func() any {
			return func(c any) func(any) func() any {
				return func(d any) func() any {
					return func() any {
						return f.(func(any) any)(a, b, c, d)
					}
				}
			}
		}
	}
}

func RunEffectFn5(f any) func(any) func(any) func(any) func(any) func(any) func() any {
	return func(a any) func(any) func(any) func(any) func(any) func() any {
		return func(b any) func(any) func(any) func(any) func() any {
			return func(c any) func(any) func(any) func() any {
				return func(d any) func(any) func() any {
					return func(e any) func() any {
						return func() any {
							return f.(func(any) any)(a, b, c, d, e)
						}
					}
				}
			}
		}
	}
}

func RunEffectFn6(f any) func(any) func(any) func(any) func(any) func(any) func(any) func() any {
	return func(a any) func(any) func(any) func(any) func(any) func(any) func() any {
		return func(b any) func(any) func(any) func(any) func(any) func() any {
			return func(c any) func(any) func(any) func(any) func() any {
				return func(d any) func(any) func(any) func() any {
					return func(e any) func(any) func() any {
						return func(g any) func() any {
							return func() any {
								var args []any
								args = append(args, a, b, c, d, e, g)
								return f.(func(any) any)(args)
							}
						}
					}
				}
			}
		}
	}
}

func RunEffectFn7(f any) func(any) func(any) func(any) func(any) func(any) func(any) func(any) func() any {
	return func(a any) func(any) func(any) func(any) func(any) func(any) func(any) func() any {
		return func(b any) func(any) func(any) func(any) func(any) func(any) func() any {
			return func(c any) func(any) func(any) func(any) func(any) func() any {
				return func(d any) func(any) func(any) func(any) func() any {
					return func(e any) func(any) func(any) func() any {
						return func(g any) func(any) func() any {
							return func(h any) func() any {
								return func() any {
									var args []any
									args = append(args, a, b, c, d, e, g, h)
									return f.(func(any) any)(args)
								}
							}
						}
					}
				}
			}
		}
	}
}

func RunEffectFn8(f any) func(any) func(any) func(any) func(any) func(any) func(any) func(any) func(any) func() any {
	return func(a any) func(any) func(any) func(any) func(any) func(any) func(any) func(any) func() any {
		return func(b any) func(any) func(any) func(any) func(any) func(any) func(any) func() any {
			return func(c any) func(any) func(any) func(any) func(any) func(any) func() any {
				return func(d any) func(any) func(any) func(any) func(any) func() any {
					return func(e any) func(any) func(any) func(any) func() any {
						return func(g any) func(any) func(any) func() any {
							return func(h any) func(any) func() any {
								return func(i any) func() any {
									return func() any {
										var args []any
										args = append(args, a, b, c, d, e, g, h, i)
										return f.(func(any) any)(args)
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

func RunEffectFn9(f any) func(any) func(any) func(any) func(any) func(any) func(any) func(any) func(any) func(any) func() any {
	return func(a any) func(any) func(any) func(any) func(any) func(any) func(any) func(any) func(any) func() any {
		return func(b any) func(any) func(any) func(any) func(any) func(any) func(any) func(any) func() any {
			return func(c any) func(any) func(any) func(any) func(any) func(any) func(any) func() any {
				return func(d any) func(any) func(any) func(any) func(any) func(any) func() any {
					return func(e any) func(any) func(any) func(any) func(any) func() any {
						return func(g any) func(any) func(any) func(any) func() any {
							return func(h any) func(any) func(any) func() any {
								return func(i any) func(any) func() any {
									return func(j any) func() any {
										return func() any {
											var args []any
											args = append(args, a, b, c, d, e, g, h, i, j)
											return f.(func(any) any)(args)
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

func RunEffectFn10(f any) func(any) func(any) func(any) func(any) func(any) func(any) func(any) func(any) func(any) func(any) func() any {
	return func(a any) func(any) func(any) func(any) func(any) func(any) func(any) func(any) func(any) func(any) func() any {
		return func(b any) func(any) func(any) func(any) func(any) func(any) func(any) func(any) func(any) func() any {
			return func(c any) func(any) func(any) func(any) func(any) func(any) func(any) func(any) func() any {
				return func(d any) func(any) func(any) func(any) func(any) func(any) func(any) func() any {
					return func(e any) func(any) func(any) func(any) func(any) func(any) func() any {
						return func(g any) func(any) func(any) func(any) func(any) func() any {
							return func(h any) func(any) func(any) func(any) func() any {
								return func(i any) func(any) func(any) func() any {
									return func(j any) func(any) func() any {
										return func(k any) func() any {
											return func() any {
												var args []any
												args = append(args, a, b, c, d, e, g, h, i, j, k)
												return f.(func(any) any)(args)
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
func Call_mkEffectFn1(arg0 any) func(any) func() any {
	return MkEffectFn1(arg0)
}
var _Gopurs_MkEffectFn1 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MkEffectFn1(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res(arg.PtrVal())
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res()
			return gopurs_runtime.Box(inner_res)
		})
		})
})
func Call_mkEffectFn2(arg0 any) func(any, any) func() any {
	return MkEffectFn2(arg0)
}
var _Gopurs_MkEffectFn2 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MkEffectFn2(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res(arg.PtrVal())
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res()
			return gopurs_runtime.Box(inner_res)
		})
		})
})
func Call_mkEffectFn3(arg0 any) func(any, any, any) func() any {
	return MkEffectFn3(arg0)
}
var _Gopurs_MkEffectFn3 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MkEffectFn3(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res(arg.PtrVal())
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res()
			return gopurs_runtime.Box(inner_res)
		})
		})
})
func Call_mkEffectFn4(arg0 any) func(any, any, any, any) func() any {
	return MkEffectFn4(arg0)
}
var _Gopurs_MkEffectFn4 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MkEffectFn4(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res(arg.PtrVal())
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res()
			return gopurs_runtime.Box(inner_res)
		})
		})
})
func Call_mkEffectFn5(arg0 any) func(any, any, any, any, any) func() any {
	return MkEffectFn5(arg0)
}
var _Gopurs_MkEffectFn5 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MkEffectFn5(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res(arg.PtrVal())
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res()
			return gopurs_runtime.Box(inner_res)
		})
		})
})
func Call_mkEffectFn6(arg0 any) func(any, any, any, any, any, any) func() any {
	return MkEffectFn6(arg0)
}
var _Gopurs_MkEffectFn6 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MkEffectFn6(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res(arg.PtrVal())
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res()
			return gopurs_runtime.Box(inner_res)
		})
		})
})
func Call_mkEffectFn7(arg0 any) func(any, any, any, any, any, any, any) func() any {
	return MkEffectFn7(arg0)
}
var _Gopurs_MkEffectFn7 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MkEffectFn7(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res(arg.PtrVal())
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res()
			return gopurs_runtime.Box(inner_res)
		})
		})
})
func Call_mkEffectFn8(arg0 any) func(any, any, any, any, any, any, any, any) func() any {
	return MkEffectFn8(arg0)
}
var _Gopurs_MkEffectFn8 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MkEffectFn8(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res(arg.PtrVal())
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res()
			return gopurs_runtime.Box(inner_res)
		})
		})
})
func Call_mkEffectFn9(arg0 any) func(any, any, any, any, any, any, any, any, any) func() any {
	return MkEffectFn9(arg0)
}
var _Gopurs_MkEffectFn9 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MkEffectFn9(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res(arg.PtrVal())
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res()
			return gopurs_runtime.Box(inner_res)
		})
		})
})
func Call_mkEffectFn10(arg0 any) func(any, any, any, any, any, any, any, any, any, any) func() any {
	return MkEffectFn10(arg0)
}
var _Gopurs_MkEffectFn10 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MkEffectFn10(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res(arg.PtrVal())
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res()
			return gopurs_runtime.Box(inner_res)
		})
		})
})
func Call_runEffectFn1(arg0 any) func(any) func() any {
	return RunEffectFn1(arg0)
}
var _Gopurs_RunEffectFn1 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := RunEffectFn1(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res(arg.PtrVal())
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res()
			return gopurs_runtime.Box(inner_res)
		})
		})
})
func Call_runEffectFn2(arg0 any) func(any) func(any) func() any {
	return RunEffectFn2(arg0)
}
var _Gopurs_RunEffectFn2 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := RunEffectFn2(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res(arg.PtrVal())
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg.PtrVal())
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res()
			return gopurs_runtime.Box(inner_res)
		})
		})
		})
})
func Call_runEffectFn3(arg0 any) func(any) func(any) func(any) func() any {
	return RunEffectFn3(arg0)
}
var _Gopurs_RunEffectFn3 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := RunEffectFn3(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res(arg.PtrVal())
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg.PtrVal())
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg.PtrVal())
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res()
			return gopurs_runtime.Box(inner_res)
		})
		})
		})
		})
})
func Call_runEffectFn4(arg0 any) func(any) func(any) func(any) func(any) func() any {
	return RunEffectFn4(arg0)
}
var _Gopurs_RunEffectFn4 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := RunEffectFn4(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res(arg.PtrVal())
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg.PtrVal())
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg.PtrVal())
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg.PtrVal())
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res()
			return gopurs_runtime.Box(inner_res)
		})
		})
		})
		})
		})
})
func Call_runEffectFn5(arg0 any) func(any) func(any) func(any) func(any) func(any) func() any {
	return RunEffectFn5(arg0)
}
var _Gopurs_RunEffectFn5 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := RunEffectFn5(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res(arg.PtrVal())
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg.PtrVal())
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg.PtrVal())
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg.PtrVal())
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg.PtrVal())
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res()
			return gopurs_runtime.Box(inner_res)
		})
		})
		})
		})
		})
		})
})
func Call_runEffectFn6(arg0 any) func(any) func(any) func(any) func(any) func(any) func(any) func() any {
	return RunEffectFn6(arg0)
}
var _Gopurs_RunEffectFn6 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := RunEffectFn6(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res(arg.PtrVal())
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg.PtrVal())
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg.PtrVal())
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg.PtrVal())
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg.PtrVal())
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg.PtrVal())
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res()
			return gopurs_runtime.Box(inner_res)
		})
		})
		})
		})
		})
		})
		})
})
func Call_runEffectFn7(arg0 any) func(any) func(any) func(any) func(any) func(any) func(any) func(any) func() any {
	return RunEffectFn7(arg0)
}
var _Gopurs_RunEffectFn7 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := RunEffectFn7(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res(arg.PtrVal())
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg.PtrVal())
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg.PtrVal())
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg.PtrVal())
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg.PtrVal())
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg.PtrVal())
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg.PtrVal())
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res()
			return gopurs_runtime.Box(inner_res)
		})
		})
		})
		})
		})
		})
		})
		})
})
func Call_runEffectFn8(arg0 any) func(any) func(any) func(any) func(any) func(any) func(any) func(any) func(any) func() any {
	return RunEffectFn8(arg0)
}
var _Gopurs_RunEffectFn8 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := RunEffectFn8(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res(arg.PtrVal())
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg.PtrVal())
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg.PtrVal())
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg.PtrVal())
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg.PtrVal())
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg.PtrVal())
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg.PtrVal())
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg.PtrVal())
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res()
			return gopurs_runtime.Box(inner_res)
		})
		})
		})
		})
		})
		})
		})
		})
		})
})
func Call_runEffectFn9(arg0 any) func(any) func(any) func(any) func(any) func(any) func(any) func(any) func(any) func(any) func() any {
	return RunEffectFn9(arg0)
}
var _Gopurs_RunEffectFn9 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := RunEffectFn9(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res(arg.PtrVal())
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg.PtrVal())
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg.PtrVal())
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg.PtrVal())
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg.PtrVal())
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg.PtrVal())
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg.PtrVal())
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg.PtrVal())
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg.PtrVal())
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res()
			return gopurs_runtime.Box(inner_res)
		})
		})
		})
		})
		})
		})
		})
		})
		})
		})
})
func Call_runEffectFn10(arg0 any) func(any) func(any) func(any) func(any) func(any) func(any) func(any) func(any) func(any) func(any) func() any {
	return RunEffectFn10(arg0)
}
var _Gopurs_RunEffectFn10 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := RunEffectFn10(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res(arg.PtrVal())
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg.PtrVal())
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg.PtrVal())
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg.PtrVal())
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg.PtrVal())
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg.PtrVal())
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg.PtrVal())
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg.PtrVal())
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg.PtrVal())
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg.PtrVal())
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res()
			return gopurs_runtime.Box(inner_res)
		})
		})
		})
		})
		})
		})
		})
		})
		})
		})
		})
})
