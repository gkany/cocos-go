function param1(str)
    chainhelper:log("param1 test. str:"..str)
end

function param2(str, flag)
    if (flag)
    then
        chainhelper:log("param2 test. flag true. str:"..str)
    else
        chainhelper:log("param2 test. flag false. str:"..str)
    end
end