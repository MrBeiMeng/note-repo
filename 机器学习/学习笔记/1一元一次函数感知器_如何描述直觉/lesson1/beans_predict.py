import dataset
from matplotlib import pyplot as plt

# xs表示豆豆的尺寸
# ys表示豆豆的毒性
# 本demo的目的是推测出一个函数，使其可以通过xs预测ys的值

xs,ys=dataset.get_beans(100)
print(xs)
print(ys)

# 坐标系

# 配置图像
plt.title("Size-Toxicity Function",fontsize=12)
plt.xlabel("Bean Size")# 设置横坐标的名字
plt.ylabel("Toxicity") # 设置纵坐标的名字
# plt.show()


# 散点图
plt.scatter(xs,ys)

# 预测函数
# y = 0.5 * x
w = 0.5
y_pre = w * xs # numpy 中的array类型可以自动解析乘法，广播到每个元素
plt.plot(xs,y_pre) # 绘制预测曲线


plt.show()
