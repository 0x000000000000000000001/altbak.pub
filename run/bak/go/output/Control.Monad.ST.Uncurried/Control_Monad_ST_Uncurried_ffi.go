package Control_Monad_ST_Uncurried

import "gopurs/output/gopurs_runtime"

func MkSTFn1(fn interface{}) interface{} {
	return func(a interface{}) interface{} {
		return fn.(func(interface{}) interface{})(a).(func(interface{}) interface{})(nil)
	}
}

func MkSTFn2(fn interface{}) interface{} {
	return func(a interface{}) interface{} {
		return func(b interface{}) interface{} {
			return fn.(func(interface{}) interface{})(a).(func(interface{}) interface{})(b).(func(interface{}) interface{})(nil)
		}
	}
}

func MkSTFn3(fn interface{}) interface{} {
	return func(a interface{}) interface{} {
		return func(b interface{}) interface{} {
			return func(c interface{}) interface{} {
				return fn.(func(interface{}) interface{})(a).(func(interface{}) interface{})(b).(func(interface{}) interface{})(c).(func(interface{}) interface{})(nil)
			}
		}
	}
}

func MkSTFn4(fn interface{}) interface{} {
	return func(a interface{}) interface{} {
		return func(b interface{}) interface{} {
			return func(c interface{}) interface{} {
				return func(d interface{}) interface{} {
					return fn.(func(interface{}) interface{})(a).(func(interface{}) interface{})(b).(func(interface{}) interface{})(c).(func(interface{}) interface{})(d).(func(interface{}) interface{})(nil)
				}
			}
		}
	}
}

func MkSTFn5(fn interface{}) interface{} {
	return func(a interface{}) interface{} {
		return func(b interface{}) interface{} {
			return func(c interface{}) interface{} {
				return func(d interface{}) interface{} {
					return func(e interface{}) interface{} {
						return fn.(func(interface{}) interface{})(a).(func(interface{}) interface{})(b).(func(interface{}) interface{})(c).(func(interface{}) interface{})(d).(func(interface{}) interface{})(e).(func(interface{}) interface{})(nil)
					}
				}
			}
		}
	}
}

func MkSTFn6(fn interface{}) interface{} {
	return func(a interface{}) interface{} {
		return func(b interface{}) interface{} {
			return func(c interface{}) interface{} {
				return func(d interface{}) interface{} {
					return func(e interface{}) interface{} {
						return func(f interface{}) interface{} {
							return fn.(func(interface{}) interface{})(a).(func(interface{}) interface{})(b).(func(interface{}) interface{})(c).(func(interface{}) interface{})(d).(func(interface{}) interface{})(e).(func(interface{}) interface{})(f).(func(interface{}) interface{})(nil)
						}
					}
				}
			}
		}
	}
}

func MkSTFn7(fn interface{}) interface{} {
	return func(a interface{}) interface{} {
		return func(b interface{}) interface{} {
			return func(c interface{}) interface{} {
				return func(d interface{}) interface{} {
					return func(e interface{}) interface{} {
						return func(f interface{}) interface{} {
							return func(g interface{}) interface{} {
								return fn.(func(interface{}) interface{})(a).(func(interface{}) interface{})(b).(func(interface{}) interface{})(c).(func(interface{}) interface{})(d).(func(interface{}) interface{})(e).(func(interface{}) interface{})(f).(func(interface{}) interface{})(g).(func(interface{}) interface{})(nil)
							}
						}
					}
				}
			}
		}
	}
}

func MkSTFn8(fn interface{}) interface{} {
	return func(a interface{}) interface{} {
		return func(b interface{}) interface{} {
			return func(c interface{}) interface{} {
				return func(d interface{}) interface{} {
					return func(e interface{}) interface{} {
						return func(f interface{}) interface{} {
							return func(g interface{}) interface{} {
								return func(h interface{}) interface{} {
									return fn.(func(interface{}) interface{})(a).(func(interface{}) interface{})(b).(func(interface{}) interface{})(c).(func(interface{}) interface{})(d).(func(interface{}) interface{})(e).(func(interface{}) interface{})(f).(func(interface{}) interface{})(g).(func(interface{}) interface{})(h).(func(interface{}) interface{})(nil)
								}
							}
						}
					}
				}
			}
		}
	}
}

func MkSTFn9(fn interface{}) interface{} {
	return func(a interface{}) interface{} {
		return func(b interface{}) interface{} {
			return func(c interface{}) interface{} {
				return func(d interface{}) interface{} {
					return func(e interface{}) interface{} {
						return func(f interface{}) interface{} {
							return func(g interface{}) interface{} {
								return func(h interface{}) interface{} {
									return func(i interface{}) interface{} {
										return fn.(func(interface{}) interface{})(a).(func(interface{}) interface{})(b).(func(interface{}) interface{})(c).(func(interface{}) interface{})(d).(func(interface{}) interface{})(e).(func(interface{}) interface{})(f).(func(interface{}) interface{})(g).(func(interface{}) interface{})(h).(func(interface{}) interface{})(i).(func(interface{}) interface{})(nil)
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

func MkSTFn10(fn interface{}) interface{} {
	return func(a interface{}) interface{} {
		return func(b interface{}) interface{} {
			return func(c interface{}) interface{} {
				return func(d interface{}) interface{} {
					return func(e interface{}) interface{} {
						return func(f interface{}) interface{} {
							return func(g interface{}) interface{} {
								return func(h interface{}) interface{} {
									return func(i interface{}) interface{} {
										return func(j interface{}) interface{} {
											return fn.(func(interface{}) interface{})(a).(func(interface{}) interface{})(b).(func(interface{}) interface{})(c).(func(interface{}) interface{})(d).(func(interface{}) interface{})(e).(func(interface{}) interface{})(f).(func(interface{}) interface{})(g).(func(interface{}) interface{})(h).(func(interface{}) interface{})(i).(func(interface{}) interface{})(j).(func(interface{}) interface{})(nil)
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

func RunSTFn1(fn interface{}) interface{} {
	return func(a interface{}) interface{} {
		return func(_ interface{}) interface{} {
			return fn.(func(interface{}) interface{})(a)
		}
	}
}

func RunSTFn2(fn interface{}) interface{} {
	return func(a interface{}) interface{} {
		return func(b interface{}) interface{} {
			return func(_ interface{}) interface{} {
				return fn.(func(interface{}) interface{})(a).(func(interface{}) interface{})(b)
			}
		}
	}
}

func RunSTFn3(fn interface{}) interface{} {
	return func(a interface{}) interface{} {
		return func(b interface{}) interface{} {
			return func(c interface{}) interface{} {
				return func(_ interface{}) interface{} {
					return fn.(func(interface{}) interface{})(a).(func(interface{}) interface{})(b).(func(interface{}) interface{})(c)
				}
			}
		}
	}
}

func RunSTFn4(fn interface{}) interface{} {
	return func(a interface{}) interface{} {
		return func(b interface{}) interface{} {
			return func(c interface{}) interface{} {
				return func(d interface{}) interface{} {
					return func(_ interface{}) interface{} {
						return fn.(func(interface{}) interface{})(a).(func(interface{}) interface{})(b).(func(interface{}) interface{})(c).(func(interface{}) interface{})(d)
					}
				}
			}
		}
	}
}

func RunSTFn5(fn interface{}) interface{} {
	return func(a interface{}) interface{} {
		return func(b interface{}) interface{} {
			return func(c interface{}) interface{} {
				return func(d interface{}) interface{} {
					return func(e interface{}) interface{} {
						return func(_ interface{}) interface{} {
							return fn.(func(interface{}) interface{})(a).(func(interface{}) interface{})(b).(func(interface{}) interface{})(c).(func(interface{}) interface{})(d).(func(interface{}) interface{})(e)
						}
					}
				}
			}
		}
	}
}

func RunSTFn6(fn interface{}) interface{} {
	return func(a interface{}) interface{} {
		return func(b interface{}) interface{} {
			return func(c interface{}) interface{} {
				return func(d interface{}) interface{} {
					return func(e interface{}) interface{} {
						return func(f interface{}) interface{} {
							return func(_ interface{}) interface{} {
								return fn.(func(interface{}) interface{})(a).(func(interface{}) interface{})(b).(func(interface{}) interface{})(c).(func(interface{}) interface{})(d).(func(interface{}) interface{})(e).(func(interface{}) interface{})(f)
							}
						}
					}
				}
			}
		}
	}
}

func RunSTFn7(fn interface{}) interface{} {
	return func(a interface{}) interface{} {
		return func(b interface{}) interface{} {
			return func(c interface{}) interface{} {
				return func(d interface{}) interface{} {
					return func(e interface{}) interface{} {
						return func(f interface{}) interface{} {
							return func(g interface{}) interface{} {
								return func(_ interface{}) interface{} {
									return fn.(func(interface{}) interface{})(a).(func(interface{}) interface{})(b).(func(interface{}) interface{})(c).(func(interface{}) interface{})(d).(func(interface{}) interface{})(e).(func(interface{}) interface{})(f).(func(interface{}) interface{})(g)
								}
							}
						}
					}
				}
			}
		}
	}
}

func RunSTFn8(fn interface{}) interface{} {
	return func(a interface{}) interface{} {
		return func(b interface{}) interface{} {
			return func(c interface{}) interface{} {
				return func(d interface{}) interface{} {
					return func(e interface{}) interface{} {
						return func(f interface{}) interface{} {
							return func(g interface{}) interface{} {
								return func(h interface{}) interface{} {
									return func(_ interface{}) interface{} {
										return fn.(func(interface{}) interface{})(a).(func(interface{}) interface{})(b).(func(interface{}) interface{})(c).(func(interface{}) interface{})(d).(func(interface{}) interface{})(e).(func(interface{}) interface{})(f).(func(interface{}) interface{})(g).(func(interface{}) interface{})(h)
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

func RunSTFn9(fn interface{}) interface{} {
	return func(a interface{}) interface{} {
		return func(b interface{}) interface{} {
			return func(c interface{}) interface{} {
				return func(d interface{}) interface{} {
					return func(e interface{}) interface{} {
						return func(f interface{}) interface{} {
							return func(g interface{}) interface{} {
								return func(h interface{}) interface{} {
									return func(i interface{}) interface{} {
										return func(_ interface{}) interface{} {
											return fn.(func(interface{}) interface{})(a).(func(interface{}) interface{})(b).(func(interface{}) interface{})(c).(func(interface{}) interface{})(d).(func(interface{}) interface{})(e).(func(interface{}) interface{})(f).(func(interface{}) interface{})(g).(func(interface{}) interface{})(h).(func(interface{}) interface{})(i)
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

func RunSTFn10(fn interface{}) interface{} {
	return func(a interface{}) interface{} {
		return func(b interface{}) interface{} {
			return func(c interface{}) interface{} {
				return func(d interface{}) interface{} {
					return func(e interface{}) interface{} {
						return func(f interface{}) interface{} {
							return func(g interface{}) interface{} {
								return func(h interface{}) interface{} {
									return func(i interface{}) interface{} {
										return func(j interface{}) interface{} {
											return func(_ interface{}) interface{} {
												return fn.(func(interface{}) interface{})(a).(func(interface{}) interface{})(b).(func(interface{}) interface{})(c).(func(interface{}) interface{})(d).(func(interface{}) interface{})(e).(func(interface{}) interface{})(f).(func(interface{}) interface{})(g).(func(interface{}) interface{})(h).(func(interface{}) interface{})(i).(func(interface{}) interface{})(j)
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
func Call_mkSTFn1(arg0 interface{}) interface{} {
	return MkSTFn1(arg0)
}
var _Gopurs_MkSTFn1 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MkSTFn1(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_mkSTFn2(arg0 interface{}) interface{} {
	return MkSTFn2(arg0)
}
var _Gopurs_MkSTFn2 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MkSTFn2(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_mkSTFn3(arg0 interface{}) interface{} {
	return MkSTFn3(arg0)
}
var _Gopurs_MkSTFn3 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MkSTFn3(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_mkSTFn4(arg0 interface{}) interface{} {
	return MkSTFn4(arg0)
}
var _Gopurs_MkSTFn4 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MkSTFn4(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_mkSTFn5(arg0 interface{}) interface{} {
	return MkSTFn5(arg0)
}
var _Gopurs_MkSTFn5 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MkSTFn5(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_mkSTFn6(arg0 interface{}) interface{} {
	return MkSTFn6(arg0)
}
var _Gopurs_MkSTFn6 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MkSTFn6(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_mkSTFn7(arg0 interface{}) interface{} {
	return MkSTFn7(arg0)
}
var _Gopurs_MkSTFn7 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MkSTFn7(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_mkSTFn8(arg0 interface{}) interface{} {
	return MkSTFn8(arg0)
}
var _Gopurs_MkSTFn8 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MkSTFn8(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_mkSTFn9(arg0 interface{}) interface{} {
	return MkSTFn9(arg0)
}
var _Gopurs_MkSTFn9 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MkSTFn9(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_mkSTFn10(arg0 interface{}) interface{} {
	return MkSTFn10(arg0)
}
var _Gopurs_MkSTFn10 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MkSTFn10(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_runSTFn1(arg0 interface{}) interface{} {
	return RunSTFn1(arg0)
}
var _Gopurs_RunSTFn1 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := RunSTFn1(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_runSTFn2(arg0 interface{}) interface{} {
	return RunSTFn2(arg0)
}
var _Gopurs_RunSTFn2 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := RunSTFn2(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_runSTFn3(arg0 interface{}) interface{} {
	return RunSTFn3(arg0)
}
var _Gopurs_RunSTFn3 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := RunSTFn3(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_runSTFn4(arg0 interface{}) interface{} {
	return RunSTFn4(arg0)
}
var _Gopurs_RunSTFn4 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := RunSTFn4(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_runSTFn5(arg0 interface{}) interface{} {
	return RunSTFn5(arg0)
}
var _Gopurs_RunSTFn5 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := RunSTFn5(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_runSTFn6(arg0 interface{}) interface{} {
	return RunSTFn6(arg0)
}
var _Gopurs_RunSTFn6 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := RunSTFn6(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_runSTFn7(arg0 interface{}) interface{} {
	return RunSTFn7(arg0)
}
var _Gopurs_RunSTFn7 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := RunSTFn7(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_runSTFn8(arg0 interface{}) interface{} {
	return RunSTFn8(arg0)
}
var _Gopurs_RunSTFn8 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := RunSTFn8(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_runSTFn9(arg0 interface{}) interface{} {
	return RunSTFn9(arg0)
}
var _Gopurs_RunSTFn9 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := RunSTFn9(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_runSTFn10(arg0 interface{}) interface{} {
	return RunSTFn10(arg0)
}
var _Gopurs_RunSTFn10 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := RunSTFn10(go_arg0)
	return gopurs_runtime.Box(go_res)
})
