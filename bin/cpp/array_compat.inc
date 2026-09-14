exports["rangeImpl"] = boxed(std::make_shared<std::function<boxed(const boxed&, const boxed&)>>(
    [](const boxed& first, const boxed& last) -> boxed {
        const int start = unbox<int>(first), end = unbox<int>(last);
        const int step = start > end ? -1 : 1;
        array_t result;
        for (int i = start; i != end; i += step) result.emplace_back(i);
        result.emplace_back(end);
        return result;
    }));
exports["filterImpl"] = boxed(std::make_shared<std::function<boxed(const boxed&, const boxed&)>>(
    [](const boxed& predicate, const boxed& input) -> boxed {
        array_t result;
        for (const auto& value : unbox<array_t>(input))
            if (unbox<bool>(predicate(value))) result.emplace_back(value);
        return result;
    }));
