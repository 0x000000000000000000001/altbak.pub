package Effect_Uncurried

import "gopurs/output/gopurs_runtime"


func MkEffectFn1(f interface{}) func(interface{}) func() interface{} {
	return func(a interface{}) func() interface{} {
		return func() interface{} {
			return f.(func(interface{}) interface{})(a)
		}
	}
}

func MkEffectFn2(f interface{}) func(interface{}, interface{}) func() interface{} {
	return func(a, b interface{}) func() interface{} {
		return func() interface{} {
			return f.(func(interface{}) interface{})(a).(func(interface{}) interface{})(b)
		}
	}
}

func MkEffectFn3(f interface{}) func(interface{}, interface{}, interface{}) func() interface{} {
	return func(a, b, c interface{}) func() interface{} {
		return func() interface{} {
			return f.(func(interface{}) interface{})(a).(func(interface{}) interface{})(b).(func(interface{}) interface{})(c)
		}
	}
}

func MkEffectFn4(f interface{}) func(interface{}, interface{}, interface{}, interface{}) func() interface{} {
	return func(a, b, c, d interface{}) func() interface{} {
		return func() interface{} {
			return f.(func(interface{}) interface{})(a).(func(interface{}) interface{})(b).(func(interface{}) interface{})(c).(func(interface{}) interface{})(d)
		}
	}
}

func MkEffectFn5(f interface{}) func(interface{}, interface{}, interface{}, interface{}, interface{}) func() interface{} {
	return func(a, b, c, d, e interface{}) func() interface{} {
		return func() interface{} {
			return f.(func(interface{}) interface{})(a).(func(interface{}) interface{})(b).(func(interface{}) interface{})(c).(func(interface{}) interface{})(d).(func(interface{}) interface{})(e)
		}
	}
}

func MkEffectFn6(f interface{}) func(interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) func() interface{} {
	return func(a, b, c, d, e, g interface{}) func() interface{} {
		return func() interface{} {
			return f.(func(interface{}) interface{})(a).(func(interface{}) interface{})(b).(func(interface{}) interface{})(c).(func(interface{}) interface{})(d).(func(interface{}) interface{})(e).(func(interface{}) interface{})(g)
		}
	}
}

func MkEffectFn7(f interface{}) func(interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) func() interface{} {
	return func(a, b, c, d, e, g, h interface{}) func() interface{} {
		return func() interface{} {
			return f.(func(interface{}) interface{})(a).(func(interface{}) interface{})(b).(func(interface{}) interface{})(c).(func(interface{}) interface{})(d).(func(interface{}) interface{})(e).(func(interface{}) interface{})(g).(func(interface{}) interface{})(h)
		}
	}
}

func MkEffectFn8(f interface{}) func(interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) func() interface{} {
	return func(a, b, c, d, e, g, h, i interface{}) func() interface{} {
		return func() interface{} {
			return f.(func(interface{}) interface{})(a).(func(interface{}) interface{})(b).(func(interface{}) interface{})(c).(func(interface{}) interface{})(d).(func(interface{}) interface{})(e).(func(interface{}) interface{})(g).(func(interface{}) interface{})(h).(func(interface{}) interface{})(i)
		}
	}
}

func MkEffectFn9(f interface{}) func(interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) func() interface{} {
	return func(a, b, c, d, e, g, h, i, j interface{}) func() interface{} {
		return func() interface{} {
			return f.(func(interface{}) interface{})(a).(func(interface{}) interface{})(b).(func(interface{}) interface{})(c).(func(interface{}) interface{})(d).(func(interface{}) interface{})(e).(func(interface{}) interface{})(g).(func(interface{}) interface{})(h).(func(interface{}) interface{})(i).(func(interface{}) interface{})(j)
		}
	}
}

func MkEffectFn10(f interface{}) func(interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) func() interface{} {
	return func(a, b, c, d, e, g, h, i, j, k interface{}) func() interface{} {
		return func() interface{} {
			return f.(func(interface{}) interface{})(a).(func(interface{}) interface{})(b).(func(interface{}) interface{})(c).(func(interface{}) interface{})(d).(func(interface{}) interface{})(e).(func(interface{}) interface{})(g).(func(interface{}) interface{})(h).(func(interface{}) interface{})(i).(func(interface{}) interface{})(j).(func(interface{}) interface{})(k)
		}
	}
}

func RunEffectFn1(f interface{}) func(interface{}) func() interface{} {
	return func(a interface{}) func() interface{} {
		return func() interface{} {
			return f.(func(interface{}) interface{})(a)
		}
	}
}

func RunEffectFn2(f interface{}) func(interface{}) func(interface{}) func() interface{} {
	return func(a interface{}) func(interface{}) func() interface{} {
		return func(b interface{}) func() interface{} {
			return func() interface{} {
				return f.(func(interface{}) interface{})(a, b)
			}
		}
	}
}

func RunEffectFn3(f interface{}) func(interface{}) func(interface{}) func(interface{}) func() interface{} {
	return func(a interface{}) func(interface{}) func(interface{}) func() interface{} {
		return func(b interface{}) func(interface{}) func() interface{} {
			return func(c interface{}) func() interface{} {
				return func() interface{} {
					return f.(func(interface{}) interface{})(a, b, c)
				}
			}
		}
	}
}

func RunEffectFn4(f interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func() interface{} {
	return func(a interface{}) func(interface{}) func(interface{}) func(interface{}) func() interface{} {
		return func(b interface{}) func(interface{}) func(interface{}) func() interface{} {
			return func(c interface{}) func(interface{}) func() interface{} {
				return func(d interface{}) func() interface{} {
					return func() interface{} {
						return f.(func(interface{}) interface{})(a, b, c, d)
					}
				}
			}
		}
	}
}

func RunEffectFn5(f interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func() interface{} {
	return func(a interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func() interface{} {
		return func(b interface{}) func(interface{}) func(interface{}) func(interface{}) func() interface{} {
			return func(c interface{}) func(interface{}) func(interface{}) func() interface{} {
				return func(d interface{}) func(interface{}) func() interface{} {
					return func(e interface{}) func() interface{} {
						return func() interface{} {
							return f.(func(interface{}) interface{})(a, b, c, d, e)
						}
					}
				}
			}
		}
	}
}

func RunEffectFn6(f interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func() interface{} {
	return func(a interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func() interface{} {
		return func(b interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func() interface{} {
			return func(c interface{}) func(interface{}) func(interface{}) func(interface{}) func() interface{} {
				return func(d interface{}) func(interface{}) func(interface{}) func() interface{} {
					return func(e interface{}) func(interface{}) func() interface{} {
						return func(g interface{}) func() interface{} {
							return func() interface{} {
								var args []interface{}
								args = append(args, a, b, c, d, e, g)
								return f.(func(interface{}) interface{})(args)
							}
						}
					}
				}
			}
		}
	}
}

func RunEffectFn7(f interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func() interface{} {
	return func(a interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func() interface{} {
		return func(b interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func() interface{} {
			return func(c interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func() interface{} {
				return func(d interface{}) func(interface{}) func(interface{}) func(interface{}) func() interface{} {
					return func(e interface{}) func(interface{}) func(interface{}) func() interface{} {
						return func(g interface{}) func(interface{}) func() interface{} {
							return func(h interface{}) func() interface{} {
								return func() interface{} {
									var args []interface{}
									args = append(args, a, b, c, d, e, g, h)
									return f.(func(interface{}) interface{})(args)
								}
							}
						}
					}
				}
			}
		}
	}
}

func RunEffectFn8(f interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func() interface{} {
	return func(a interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func() interface{} {
		return func(b interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func() interface{} {
			return func(c interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func() interface{} {
				return func(d interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func() interface{} {
					return func(e interface{}) func(interface{}) func(interface{}) func(interface{}) func() interface{} {
						return func(g interface{}) func(interface{}) func(interface{}) func() interface{} {
							return func(h interface{}) func(interface{}) func() interface{} {
								return func(i interface{}) func() interface{} {
									return func() interface{} {
										var args []interface{}
										args = append(args, a, b, c, d, e, g, h, i)
										return f.(func(interface{}) interface{})(args)
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

func RunEffectFn9(f interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func() interface{} {
	return func(a interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func() interface{} {
		return func(b interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func() interface{} {
			return func(c interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func() interface{} {
				return func(d interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func() interface{} {
					return func(e interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func() interface{} {
						return func(g interface{}) func(interface{}) func(interface{}) func(interface{}) func() interface{} {
							return func(h interface{}) func(interface{}) func(interface{}) func() interface{} {
								return func(i interface{}) func(interface{}) func() interface{} {
									return func(j interface{}) func() interface{} {
										return func() interface{} {
											var args []interface{}
											args = append(args, a, b, c, d, e, g, h, i, j)
											return f.(func(interface{}) interface{})(args)
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

func RunEffectFn10(f interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func() interface{} {
	return func(a interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func() interface{} {
		return func(b interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func() interface{} {
			return func(c interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func() interface{} {
				return func(d interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func() interface{} {
					return func(e interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func() interface{} {
						return func(g interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func() interface{} {
							return func(h interface{}) func(interface{}) func(interface{}) func(interface{}) func() interface{} {
								return func(i interface{}) func(interface{}) func(interface{}) func() interface{} {
									return func(j interface{}) func(interface{}) func() interface{} {
										return func(k interface{}) func() interface{} {
											return func() interface{} {
												var args []interface{}
												args = append(args, a, b, c, d, e, g, h, i, j, k)
												return f.(func(interface{}) interface{})(args)
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
func Call_mkEffectFn1(arg0 interface{}) func(interface{}) func() interface{} {
	return MkEffectFn1(arg0)
}
var _Gopurs_MkEffectFn1 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MkEffectFn1(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res(arg)
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res()
			return gopurs_runtime.Box(inner_res)
		})
		})
})
func Call_mkEffectFn2(arg0 interface{}) func(interface{}, interface{}) func() interface{} {
	return MkEffectFn2(arg0)
}
var _Gopurs_MkEffectFn2 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MkEffectFn2(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res(arg)
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res()
			return gopurs_runtime.Box(inner_res)
		})
		})
})
func Call_mkEffectFn3(arg0 interface{}) func(interface{}, interface{}, interface{}) func() interface{} {
	return MkEffectFn3(arg0)
}
var _Gopurs_MkEffectFn3 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MkEffectFn3(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res(arg)
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res()
			return gopurs_runtime.Box(inner_res)
		})
		})
})
func Call_mkEffectFn4(arg0 interface{}) func(interface{}, interface{}, interface{}, interface{}) func() interface{} {
	return MkEffectFn4(arg0)
}
var _Gopurs_MkEffectFn4 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MkEffectFn4(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res(arg)
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res()
			return gopurs_runtime.Box(inner_res)
		})
		})
})
func Call_mkEffectFn5(arg0 interface{}) func(interface{}, interface{}, interface{}, interface{}, interface{}) func() interface{} {
	return MkEffectFn5(arg0)
}
var _Gopurs_MkEffectFn5 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MkEffectFn5(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res(arg)
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res()
			return gopurs_runtime.Box(inner_res)
		})
		})
})
func Call_mkEffectFn6(arg0 interface{}) func(interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) func() interface{} {
	return MkEffectFn6(arg0)
}
var _Gopurs_MkEffectFn6 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MkEffectFn6(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res(arg)
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res()
			return gopurs_runtime.Box(inner_res)
		})
		})
})
func Call_mkEffectFn7(arg0 interface{}) func(interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) func() interface{} {
	return MkEffectFn7(arg0)
}
var _Gopurs_MkEffectFn7 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MkEffectFn7(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res(arg)
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res()
			return gopurs_runtime.Box(inner_res)
		})
		})
})
func Call_mkEffectFn8(arg0 interface{}) func(interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) func() interface{} {
	return MkEffectFn8(arg0)
}
var _Gopurs_MkEffectFn8 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MkEffectFn8(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res(arg)
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res()
			return gopurs_runtime.Box(inner_res)
		})
		})
})
func Call_mkEffectFn9(arg0 interface{}) func(interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) func() interface{} {
	return MkEffectFn9(arg0)
}
var _Gopurs_MkEffectFn9 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MkEffectFn9(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res(arg)
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res()
			return gopurs_runtime.Box(inner_res)
		})
		})
})
func Call_mkEffectFn10(arg0 interface{}) func(interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) func() interface{} {
	return MkEffectFn10(arg0)
}
var _Gopurs_MkEffectFn10 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MkEffectFn10(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res(arg)
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res()
			return gopurs_runtime.Box(inner_res)
		})
		})
})
func Call_runEffectFn1(arg0 interface{}) func(interface{}) func() interface{} {
	return RunEffectFn1(arg0)
}
var _Gopurs_RunEffectFn1 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := RunEffectFn1(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res(arg)
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res()
			return gopurs_runtime.Box(inner_res)
		})
		})
})
func Call_runEffectFn2(arg0 interface{}) func(interface{}) func(interface{}) func() interface{} {
	return RunEffectFn2(arg0)
}
var _Gopurs_RunEffectFn2 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := RunEffectFn2(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res()
			return gopurs_runtime.Box(inner_res)
		})
		})
		})
})
func Call_runEffectFn3(arg0 interface{}) func(interface{}) func(interface{}) func(interface{}) func() interface{} {
	return RunEffectFn3(arg0)
}
var _Gopurs_RunEffectFn3 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := RunEffectFn3(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res()
			return gopurs_runtime.Box(inner_res)
		})
		})
		})
		})
})
func Call_runEffectFn4(arg0 interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func() interface{} {
	return RunEffectFn4(arg0)
}
var _Gopurs_RunEffectFn4 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := RunEffectFn4(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res()
			return gopurs_runtime.Box(inner_res)
		})
		})
		})
		})
		})
})
func Call_runEffectFn5(arg0 interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func() interface{} {
	return RunEffectFn5(arg0)
}
var _Gopurs_RunEffectFn5 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := RunEffectFn5(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
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
func Call_runEffectFn6(arg0 interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func() interface{} {
	return RunEffectFn6(arg0)
}
var _Gopurs_RunEffectFn6 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := RunEffectFn6(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
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
func Call_runEffectFn7(arg0 interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func() interface{} {
	return RunEffectFn7(arg0)
}
var _Gopurs_RunEffectFn7 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := RunEffectFn7(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
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
func Call_runEffectFn8(arg0 interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func() interface{} {
	return RunEffectFn8(arg0)
}
var _Gopurs_RunEffectFn8 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := RunEffectFn8(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
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
func Call_runEffectFn9(arg0 interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func() interface{} {
	return RunEffectFn9(arg0)
}
var _Gopurs_RunEffectFn9 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := RunEffectFn9(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
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
func Call_runEffectFn10(arg0 interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func(interface{}) func() interface{} {
	return RunEffectFn10(arg0)
}
var _Gopurs_RunEffectFn10 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := RunEffectFn10(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
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
