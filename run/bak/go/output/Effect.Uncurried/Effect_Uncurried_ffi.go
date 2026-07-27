package Effect_Uncurried


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
